package db

type Music struct {
	//	Id       string `json:"_id"`
	Duration string `json:"duration"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Genre    string `json:"genre"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}
