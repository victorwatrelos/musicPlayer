package db

import (
	"gopkg.in/mgo.v2"
)

type DBManager struct {
	session         *mgo.Session
	trackCollection *mgo.Collection
}

type MusicData struct {
	Artist string
}
