package db

import (
	"../crawl"
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

func (b *DBManager) InsertData(data *crawl.Music) {
	err := b.musicCollection.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (b *DBManager) Query(label, data string) {
	var results [](crawl.Music)
	err := b.musicCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{`.*` + data + `.*`, ""}}}).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}

func (b *DBManager) Close() {
	fmt.Println("Close Mongodb session")
	b.session.Close()
}
