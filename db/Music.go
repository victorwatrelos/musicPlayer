package db

import (
	"gopkg.in/mgo.v2/bson"
)

type Track struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Duration string        `json:"duration"`
	Artist   string        `json:"artist"`
	Album    string        `json:"album"`
	Genre    string        `json:"genre"`
	Name     string        `json:"name"`
	Path     string        `json:"path"`
}
