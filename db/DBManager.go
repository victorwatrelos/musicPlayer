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

func (b *DBManager) ReadMusicChan(c chan crawl.Music) {
	for music := range c {
		b.InsertData(&music)
	}
	fmt.Println("Finishing insert music")
}

func (b *DBManager) InsertData(data *crawl.Music) {
	err := b.musicCollection.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (b *DBManager) Query(label, data string) {
	var results [](crawl.Music)
	cmd := b.musicCollection.Find(bson.M{label: bson.M{"$regex": bson.RegEx{".*.*", ""}}})
	err := cmd.All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}

func (b *DBManager) Close() {
	fmt.Println("Close Mongodb session")
	b.session.Close()
}
