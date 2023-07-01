package routes

import (
	"goupie-tracker/handler"
	"net/http"
)

func Route() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/artists", handler.GetArtists)
	http.HandleFunc("/artist/", handler.GetOneArtist)
	http.HandleFunc("/events", handler.GetEvents)
	http.HandleFunc("/event/", handler.GetOneEvent)
}
