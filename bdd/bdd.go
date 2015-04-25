package bdd

import (
	"../crawl"
	"fmt"
	"gopkg.in/mgo.v2"
)

func (b *BddManager) GetConnection() {
	fmt.Println("Opening mongodb session")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	b.session = session
	b.musicCollection = session.DB("musique").C("track")
}

func (b *BddManager) InsertData(data *crawl.Music) {
	err := b.musicCollection.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (b *BddManager) Close() {
	fmt.Println("Close Mongodb session")
	b.session.Close()
}
