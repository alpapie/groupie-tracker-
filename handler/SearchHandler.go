package handler

import (
	"encoding/json"
	"fmt"
	"goupie-tracker/helper"
	"goupie-tracker/middlewares"
	"goupie-tracker/models"
	"net/http"
	"strings"
)

type Item struct {
	Items []string `json:"Items"`
}

func Search(w http.ResponseWriter, r *http.Request) {
	_ok, pageError := middlewares.CheckRequest(r, "/search", "get")
	search := r.URL.Query().Get("search")
	if _ok {
		 Datas:=  struct {
			Artists []models.Artist
			Title   string
		}{
			Artists: nil,
			Title:"Artists",
		}

		_=RefreshData(true, false, true, false)
		if search != "" {
			// list:=strings.Split(search," - ")
			// if len(list)<2{
			// 	Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,false,false)
			// }else{
			// 	search=list[0]
			// 	if list[1]=="artist/band"{
			// 		Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,true,false,false)
			// 		} else if list[1]=="member"{
			// 		fmt.Println(list[1])
			// 		Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,true,false)
			// 	}else if list[1]=="location"{
			// 		Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,false,true)
			// 	}
			// }
			if strings.HasSuffix(search," - artist/band") {
				search=strings.ReplaceAll(search," - artist/band","")
				Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,true,false,false)
			} else if strings.HasSuffix(search," - member"){
				search=strings.ReplaceAll(search," - member","")
				Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,true,false)
			}else if strings.HasSuffix(search," - location"){
				search=strings.ReplaceAll(search," - location","")
				Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,false,true)
			}else{
				Datas.Artists=helper.FilterData(Artists,Location.Indexes,search,false,false,true)
			}
			err := helper.RenderTemplateWithLoyout(w, "pages/artist/artist", Datas)
			if err != nil {
				helper.ErrorPage(w, 404)
				fmt.Println(err)
			}
			return 
		} else {
			Datas.Artists=Artists
			err := helper.RenderTemplateWithLoyout(w, "pages/artist/artist", Datas)
			if err != nil {
				helper.ErrorPage(w, 404)
				fmt.Println(err)
			}
			return
		}
	}
	helper.ErrorPage(w, pageError)
	return
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	_ok, pageError := middlewares.CheckRequest(r, "/searchItem", "get")
	if _ok {
		_item := r.URL.Query().Get("param")
		if !RefreshData(true, false, true, false) {
			if _item == "" {
				_item = ""
			}
			Items := Item{
				Items: helper.FilterLoc(helper.GetArrayloc(Location.Indexes, Artists), _item),
			}
			// Convert data to JSON
			response, err := json.Marshal(Items)
			if err != nil {
				fmt.Println(err)
			}
			// Set the response headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Write the JSON response
			w.Write(response)
			return
		}
	}
	w.WriteHeader(pageError)
}
