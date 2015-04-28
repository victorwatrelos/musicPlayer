package server

import (
	"../db"
	"../player"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "log"
	"net/http"
	"strconv"
)

type jsonRet struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type vol_ret struct {
	Volume int `json:"volume"`
}

type choose_ret struct {
	Path string `json:"path"`
}

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
	p := player.GetSing()
	p.Play()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Play", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func Pause(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	p.Pause()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Pause", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var t choose_ret
	if err = json.Unmarshal(body, &t); err != nil {
		panic(err)
	}

	p.Add(t.Path)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Track " + t.Name + " by " + t.Artist + " To Playlist", Code: http.StatusOK}); err != nil {
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Track " + t.Path + " To Playlist", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ChVolume(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var t vol_ret
	if err = json.Unmarshal(body, &t); err != nil {
		panic(err)
	}
	p.SetVolume(t.Volume)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Volume changed to " + strconv.Itoa(t.Volume), Code: http.StatusOK}); err != nil {
		panic(err)
	}
}
