package server

import (
	"../db"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

type vol_ret struct {
	Volume int `json:"volume"`
}

type choose_ret struct {
	Path string `json:"path"`
}

func (v *vol_ret) UnmarshalHTTP(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, v)
	return err
}

func (c *choose_ret) UnmarshalHTTP(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, c)
	return err
}

func GetBody(r *http.Request, v Entity) error {
	return v.UnmarshalHTTP(r)
}

func GetData(label, query string) []db.Music {
	dtb := db.DBManager{}
	dtb.GetConnection()
	defer dtb.Close()
	results := dtb.Query(label, query)
	return results
}

func GetDataDistinct(label, query string) []string {
	dtb := db.DBManager{}
	dtb.GetConnection()
	defer dtb.Close()
	results := dtb.DistinctQ(label, label, query, label)
	return results
}

func MapQuery(r *http.Request, label string, fn func(string)) string {
	var t choose_ret
	if err := GetBody(r, &t); err != nil {
		panic(err)
	}
	results := GetData("artist", t.Path)
	for _, track := range results {
		fn(track.Path)
	}
	return t.Path
}
