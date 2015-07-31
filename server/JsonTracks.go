package server

import (
	"encoding/json"
	"io/ioutil"
	_ "log"
	"net/http"
)

type JsonTracks struct {
	Tracks []string `json:"tracks"`
}

func (jsonTracks *JsonTracks) UnmarshalHTTP(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, jsonTracks)
	return err
}
