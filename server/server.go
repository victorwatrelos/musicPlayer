package server

import (
	"log"
	"net/http"
)

func Launch() {

	router := NewRouter()
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
