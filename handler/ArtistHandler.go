package handler

import (
	"fmt"
	"goupie-tracker/helper"
	"goupie-tracker/middlewares"
	"goupie-tracker/models"
	"net/http"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	// var ArtistsAll []models.Artist
	ok, PageError := middlewares.CheckRequest(r, "/artists", "get")
	if !ok {
		helper.ErrorPage(w, PageError)
		return
	}
	if !RefreshData(true,false,false,false){
		Datas := struct {
			Artists []models.Artist
			Title       string
		}{
			Artists: Artists,
			Title:"Artists",
		}
		// fmt.Println(Datas.Artists[0])
		err := helper.RenderTemplateWithLoyout(w, "pages/artist/artist", Datas)
		if err != nil {
			helper.ErrorPage(w, 404)
			fmt.Println(err)
		}
		return 
	} 
	helper.ErrorPage(w, 500)
	return
}

func GetOneArtist(w http.ResponseWriter, r *http.Request) {
	OneArtist := models.ArtistOne{}
	_ok := middlewares.Getmethode(r, "get")
	if !_ok {
		helper.ErrorPage(w, 405)
		return
	}
	okk,id:=helper.PArseUlr(r,"artist") 
	if okk {
		ok := OneArtist.GetOneartist(id)
		if ok {
			err := helper.RenderTemplate(w, "pages/artist/artist single", OneArtist)
			if err != nil {
				helper.ErrorPage(w, 404)
				fmt.Println(err)
				return
			}
			return
		} 
	}
	helper.ErrorPage(w, 404)
	
}
