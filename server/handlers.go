package server

import (
	"../db"
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func GetData(label string, w http.ResponseWriter, r *http.Request) {
	dtb := db.DBManager{}
	dtb.GetConnection()
	defer dtb.Close()
	var results []db.Music

	q := r.URL.Query().Get("q")
	results = dtb.Query(label, q)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}

}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	GetData("genre", w, r)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	GetData("album", w, r)
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	GetData("artist", w, r)
}

func Play(w http.ResponseWriter, r *http.Request) {

}

func Pause(w http.ResponseWriter, r *http.Request) {

}

func Add(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if len(q) != 0 {

	}
}

func ChVolume(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if len(q) != 0 {

	}
}
