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
		"Show Playlist",
		"GET",
		"/playlist",
		ShowPlaylist,
	},
	Route{
		"Show Status",
		"GET",
		"/status",
		ShowStatus,
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
		Prev,
	},
	Route{
		"Next",
		"PUT",
		"/next",
		Next,
	},
	Route{
		"Change Volume",
		"PUT",
		"/vol",
		ChVolume,
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
		"/add/artist",
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
		"/queue/artist",
		AddArtistToQueue,
	},
	Route{
		"Clear Playlist Queue",
		"PUT",
		"/clear",
		ClearQueue,
	},
	Route{
		"GoTo Playlist's index",
		"PUT",
		"/goto",
		Goto,
	},
	Route{
		"ToggleLoop",
		"PUT",
		"/loop",
		SetLoop,
	},
	Route{
		"ToggleRandom",
		"PUT",
		"/random",
		SetRandom,
	},
	Route{
		"ToggleRepeat",
		"PUT",
		"/repeat",
		SetRepeat,
	},
}
