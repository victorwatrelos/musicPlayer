package server

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Show Artist",
		"GET",
		"/artist",
		GetArtist,
	},
	Route{
		"Show Album",
		"GET",
		"/artist/album",
		GetAlbum,
	},
	Route{
		"Show Genre",
		"GET",
		"/genre",
		GetGenre,
	},
	Route{
		"Play",
		"GET",
		"/play",
		Play,
	},
	Route{
		"Pause",
		"GET",
		"/pause",
		Pause,
	},
	Route{
		"Add",
		"POST",
		"/add",
		Add,
	},
	Route{
		"Volume",
		"PUT",
		"/vol",
		ChVolume,
	},
}
