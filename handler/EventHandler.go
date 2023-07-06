package handler

import (
	"fmt"
	"goupie-tracker/helper"
	"goupie-tracker/middlewares"
	"goupie-tracker/models"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	// var ArtistsAll []models.Artist
	ok, PageError := middlewares.CheckRequest(r, "/events", "get")
	if !ok {
		helper.ErrorPage(w, PageError)
		return
	}
	err := helper.GetJson(Url["artist"], &Artists)
	err1 := helper.GetJson(Url["event"], &Events)
	if err == nil && err1==nil{
		Datas := struct {
			Events []models.Relation
			Title   string
		}{
			Events:	 Events.Indexes,
			Title:   "Events",
		}
		for i := range Datas.Events {
			Datas.Events[i].Artistname=Artists[i].Name
		}
		// fmt.Println(Datas.Artists[0])
		err := helper.RenderTemplateWithLoyout(w, "pages/event/event", Datas)
		if err != nil {
			helper.ErrorPage(w, 404)
			fmt.Println(err)
			return 
		}
	} else {
		helper.ErrorPage(w, 500)
		return
	}
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	OneEvents:= models.ArtistOne{}
	_ok := middlewares.Getmethode(r, "get")
	if !_ok {
		helper.ErrorPage(w, 405)
		return
	}
	okk,id:=helper.PArseUlr(r,"event")
	loc:=r.URL.Query().Get("location")
	if okk {
		ok := OneEvents.GetOneartist(id)
		if ok {
			if loc==""{
				for i:= range OneEvents.Relations.DatesLocations{
					loc=i
					break
				}
			}
			OneEvents.Location=loc
			err := helper.RenderTemplate(w, "pages/event/concert single", OneEvents)
			if err != nil {
				helper.ErrorPage(w, 404)
				fmt.Println(err)
				return
			}
			return
		}
	}
	helper.ErrorPage(w, 404)
	return
}
