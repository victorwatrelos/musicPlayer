package bdd

import (
	"gopkg.in/mgo.v2"
)

type BddManager struct {
	session         *mgo.Session
	musicCollection *mgo.Collection
}

type MusicData struct {
	Artist string
}
