package server

import (
	"log"
	"net/http"
)

func Launch() {
	log.Println("Server Launch")
	router := NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
