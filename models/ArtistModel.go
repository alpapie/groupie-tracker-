package models

import (
	"goupie-tracker/helper"
	"strconv"
)

var Url = map[string]string{
	"artist":   "https://groupietrackers.herokuapp.com/api/artists",
	"event":    "https://groupietrackers.herokuapp.com/api/relation",
	"location": "https://groupietrackers.herokuapp.com/api/locations",
	"date":     "https://groupietrackers.herokuapp.com/api/date/",
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type ArtistOne struct {
	Artists      Artist
	Locations    Location
	ConcertDates Date
	Relations    Relation
}

//****************************** GET ARTIST METHODE ******************************************************
// func (ArtistOne) GetAllartists() (bool, []ArtistOne) {

// }

//****************************** GET ARTIST METHODE ******************************************************
func (AllArtists *ArtistOne) GetOneartist(id int) bool {
	artist := Artist{}
	var err1, err2 error

	err := helper.GetJson(Url["artist"]+"/"+strconv.Itoa(id), &artist)
	// fmt.Println(artist)
	if err == nil {
		location := Location{}
		relation := Relation{}
		// fmt.Println("alpapie")

		err1 = helper.GetJson(artist.Locations, &location)
		err2 = helper.GetJson(artist.Relations, &relation)

		if err1 == nil && err2 == nil {
			AllArtists.Artists = artist
			AllArtists.Locations = location
			AllArtists.Relations = relation
			AllArtists.Relations.Artistname = artist.Name
		} else {
			return false
		}
		return true
	} else {
		return false
	}
}
