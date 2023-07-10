package models

import (
	"encoding/json"
	"net/http"
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
	Artists   Artist
	Relations Relation
	Location  string
}

//****************************** GET ARTIST METHODE ******************************************************
// func (ArtistOne) GetAllartists() (bool, []ArtistOne) {

// }

//****************************** GET ARTIST METHODE ******************************************************
func (AllArtists *ArtistOne) GetOneartist(id int) bool {
	artist := Artist{}
	var err1, err2 error

	err := GetJson(Url["artist"]+"/"+strconv.Itoa(id), &artist)
	// fmt.Println(artist)
	if err == nil {
		relation := Relation{}
		// fmt.Println("alpapie")
		err2 = GetJson(artist.Relations, &relation)

		if err1 == nil && err2 == nil {
			AllArtists.Artists = artist
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

func GetJson(url string, model interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}
