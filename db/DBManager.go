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
	b.trackCollection = session.DB("music").C("track")
}

func (b *DBManager) ReadMusicChan(c chan Track) {
	for track := range c {
		info, err := b.trackCollection.Upsert(bson.M{"artist": track.Artist, "album": track.Album, "name": track.Name}, track)
		if err != nil {
			panic(err)
		}
		if info.UpsertedId == nil {
			fmt.Println("Update: ", track)
			continue
		}
		fmt.Println("New: ", track)
	}
	fmt.Println("Finishing insert music")
}

func (b *DBManager) FindByIds(ids []string) []Track {
	var tracks []Track
	var tracksObjectIds []bson.ObjectId

	for _, hexId := range ids {
		tracksObjectIds = append(tracksObjectIds, bson.ObjectIdHex(hexId))
	}

	b.trackCollection.Find(bson.M{"_id": bson.M{"$in": tracksObjectIds}}).All(&tracks)
	return tracks

}
func (b *DBManager) InsertData(track *Track) {
	err := b.trackCollection.Insert(track)
	if err != nil {
		panic(err)
	}
}

func (b *DBManager) Query(label, data string) []Track {
	var results []Track
	err := b.trackCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, "i"}}}).All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

func (b *DBManager) DistinctQ(label, key, data, sort string) []string {
	var results []string
	err := b.trackCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, "i"}}}).Sort(sort).Distinct(key, &results)
	if err != nil {
		panic(err)
	}
	return results
}

func (b *DBManager) Close() {
	fmt.Println("Close Mongodb session")
	b.session.Close()
}
