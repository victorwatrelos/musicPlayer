package db

import (
	"gopkg.in/mgo.v2/bson"
)

type Music struct {
	Duration string `json:"duration"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

type Track struct {
	Id       bson.ObjectId `json:"_id"`
	Duration string        `json:"duration"`
	Artist   bson.ObjectId `json:"artist"`
	Album    bson.ObjectId `json:"album"`
	Genre    bson.ObjectId `json:"genre"`
	Name     string        `json:"name"`
	Path     string        `json:"path"`
}

type Artist struct {
	Id   bson.ObjectId `bson:"_id,omitempty"	json:"id"`
	Name string        `json:"name"`
}

type Genre struct {
	Id   bson.ObjectId `bson:"_id,omitempty"	json:"id"`
	Name string        `json:"name"`
}

type Album struct {
	Id   bson.ObjectId `bson:"_id,omitempty"	json:"id"`
	Name string        `json:"name"`
}
