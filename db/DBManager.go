package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (b *DBManager) GetConnection() {
	fmt.Println("Opening mongodb session")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	b.session = session
	b.musicCollection = session.DB("musique").C("track")
	b.artistCollection = session.DB("musique").C("artist")
}

func (b *DBManager) formatData(music *Music) {
}

func (b *DBManager) ReadMusicChan(c chan Music) {
	for music := range c {
		b.formatData(&music)
		b.InsertData(&music)
	}
	fmt.Println("Finishing insert music")
}

func (b *DBManager) InsertData(music *Track, artist *Artist) {
	var artist Artist
	artist = b.SaveArtist(artist)
	err := b.musicCollection.Insert(music)
	if err != nil {
		panic(err)
	}
}

func (b *DBManager) createArtist(artist *Artist) *Artist {
	i := bson.NewObjectId()
	err := b.artistCollection.Insert(bson.M{"_id": i, "name": artist.Name})
	if err != nil {
		panic(err)
	}
	artist.Id = i
	return artist
}

func (b *DBManager) SaveArtist(artistOrig *Artist) *Artist {
	var artist Artist
	err := b.artistCollection.Find(bson.M{"name": name}).One(&artist)
	if err != nil {
		if err.Error() == "not found" {
			artist = b.createArtist(&artistOrig)
		} else {
			panic(err)
		}
	}
	return &artist
}

func (b *DBManager) Query(label, data string) []Music {
	var results []Music
	err := b.musicCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, "i"}}}).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func (b *DBManager) DistinctQ(label, key, data, sort string) []string {
	var results []string
	err := b.musicCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, "i"}}}).Sort(sort).Distinct(key, &results)
	if err != nil {
		panic(err)
	}
	return results
}

func (b *DBManager) Close() {
	fmt.Println("Close Mongodb session")
	b.session.Close()
}
