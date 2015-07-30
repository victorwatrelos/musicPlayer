package db

import ()

type Track struct {
	Duration string `json:"duration"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}
