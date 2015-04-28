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
		"Show Artists",
		"GET",
		"/artist",
		GetArtist,
	},
	Route{
		"Show Album",
		"GET",
		"/album",
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
		"PUT",
		"/play",
		Play,
	},
	Route{
		"Pause",
		"PUT",
		"/pause",
		Pause,
	},
	Route{
		"Prev",
		"PUT",
		"/prev",
		Next,
	},
	Route{
		"Next",
		"PUT",
		"/next",
		Prev,
	},
	Route{
		"Add Track To Playlist",
		"PUT",
		"/add",
		Add,
	},
	Route{
		"Add Artist To Playlist",
		"PUT",
		"/add/Artist",
		AddArtist,
	},
	Route{
		"Queue Track To Playlist",
		"PUT",
		"/queue",
		AddToQueue,
	},
	Route{
		"Queue Artist To Playlist",
		"PUT",
		"/queue/Artist",
		AddArtistToQueue,
	},
	Route{
		"Clear Playlist Queue",
		"PUT",
		"/clearQueue",
		ClearQueue,
	},
	Route{
		"Change Volume",
		"PUT",
		"/vol",
		ChVolume,
	},
}
