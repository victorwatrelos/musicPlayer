package server

import (
	"../player"
	"encoding/json"
	"fmt"
	_ "log"
	"net/http"
	"strconv"
)

type jsonRet struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//	HANDLERS

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

//	DATABASE GETTER HANDLERS

func GetGenre(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	results := GetData("genre", q)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	results := GetData("album", q)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	results := GetData("artist", q)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}

//	PLAYER INFO HANDLERS

func ShowPlaylist(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	ret := p.ShowPlaylist()
	fmt.Println(ret)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Show Playlist", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	ret := p.ShowStatus()
	fmt.Println(ret)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Show Status", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func Goto(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	p.Goto(t.Path)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Show Status", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

//	PLAYER OPTION HANDLERS

func ToggleLoop(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	if t.Path == "0" {
		p.ToggleLoop()
	} else {
		p.ToggleLoop()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Toggle Loop", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ToggleRandom(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	if t.Path == "0" {
		p.ToggleRandom()
	} else {
		p.ToggleRandom()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Toggle Random", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ToggleRepeat(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	if t.Path == "0" {
		p.ToggleRepeat()
	} else {
		p.ToggleRepeat()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Toggle Repeat", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

//	PLAYER CONTROL HANDLERS

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

func Prev(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	p.Prev()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Prev", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func Next(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	p.Next()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Next", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ChVolume(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t vol_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	p.SetVolume(t.Volume)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Volume changed to " + strconv.Itoa(t.Volume), Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

//	PLAYER'S PLAYLIST CONTROL HANDLERS

func Add(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	p.Add(t.Path)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Track " + t.Path + " To Playlist", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func AddToQueue(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	p.AddToQueue(t.Path)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Track " + t.Path + " To Playlist's Queue", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func AddArtist(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	p.ClearQueue()
	str := MapQuery(r, "artist", p.AddToQueue)
	p.Play()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Artist " + str + " To Playlist", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func AddArtistToQueue(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	str := MapQuery(r, "artist", p.AddToQueue)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Added Artist " + str + " To Playlist", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}

func ClearQueue(w http.ResponseWriter, r *http.Request) {
	p := player.GetSing()
	p.ClearQueue()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(jsonRet{Message: "Queue Cleared", Code: http.StatusOK}); err != nil {
		panic(err)
	}
}
