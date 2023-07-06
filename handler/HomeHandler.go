package handler

import (
	"fmt"
	"goupie-tracker/helper"
	"goupie-tracker/middlewares"
	"goupie-tracker/models"
	"net/http"
	"sync"
)

var Url = map[string]string{
	"artist":   "https://groupietrackers.herokuapp.com/api/artists",
	"event":    "https://groupietrackers.herokuapp.com/api/relation",
	"location": "https://groupietrackers.herokuapp.com/api/locations",
	"date":     "https://groupietrackers.herokuapp.com/api/dates",
}
var Artists []models.Artist
var Events models.Index
var Date models.IndexDate
var Location models.IndexLoc

type Data struct {
	Artists  []models.Artist
	Events   []models.Relation
	Location models.Location
	Date     models.Date
}

var A models.ArtistOne

func Index(w http.ResponseWriter, r *http.Request) {
	ok, PageError := middlewares.CheckRequest(r, "/", "get")
	var wg sync.WaitGroup
	ch :=make(chan error)

	if !ok {
		helper.ErrorPage(w, PageError)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		GetData(ch)
	}()
	

	go func() {
		wg.Wait()
		close(ch)
	}()
	if !helper.FecthError(ch) {
		Datatemp := Data{}

		Datatemp.Artists = Artists[:9]
		Datatemp.Events = Events.Indexes[:9]
		Datatemp.Date=Date.Indexes[0]
		Datatemp.Location=Location.Indexes[0]

		for i, v := range Datatemp.Artists {
			Datatemp.Events[i].Artistname = v.Name
		}
		// AllA,is:=A.GetAllartists()
		// fmt.Println(A.GetAllartists())
		err := helper.RenderTemplate(w, "pages/index", Datatemp)
		if err != nil {
			helper.ErrorPage(w, 404)
			fmt.Println(err)
			return
		}
	} else {
		// fmt.Fprintf(w, string("error to get data"),iserr)
		helper.ErrorPage(w, 500)
		return
	}
}

func GetData(ch chan error) {

	ch <- helper.GetJson(Url["artist"], &Artists)
	ch <- helper.GetJson(Url["event"], &Events)
	ch <- helper.GetJson(Url["date"], &Date)
	ch <- helper.GetJson(Url["location"], &Location)

}
