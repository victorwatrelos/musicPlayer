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
}

func (b *DBManager) ReadMusicChan(c chan Music) {
	for music := range c {
		b.InsertData(&music)
	}
	fmt.Println("Finishing insert music")
}

func (b *DBManager) InsertData(data *Music) {
	err := b.musicCollection.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (b *DBManager) Query(label, data string) []Music {
	var results []Music
	err := b.musicCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, "i"}}}).All(&results)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(results)
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
