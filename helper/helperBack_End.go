package helper

import (
	"encoding/json"
	"fmt"
	"goupie-tracker/models"
	"net/http"
	"strconv"
	"strings"
)

//******************* VERIF IF THE STRING IS AN INT*****************************************************
func IsInt(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

//***************************** GET JSON DATA ************************************************************
func GetJson(url string, model interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(model)
}

// ******************************* PARSE FILE IN URL *****************
func PArseUlr(r *http.Request, match string) (bool, int) {
	index := strings.Split(r.URL.Path[1:], "/")
	if len(index) == 2 && index[0] == match {
		id, err := strconv.Atoi(index[1])
		if err == nil {
			return true, id
		}
	}
	return false, 0
}

// *************************************** FILTER THE DATA WITH THE PARAM****************
func FilterData(Artist []models.Artist, Location []models.Location, s string, arrt, mem,loc bool) ([]models.Artist) {
	_Artist:=[]models.Artist{}
	s=strings.ToLower(s)
	for i := 0; i < len(Artist); i++ {
		if arrt{
			if strings.Contains(strings.ToLower(Artist[i].Name),s){
				_Artist=append(_Artist, Artist[i])
			}
		}else if mem && len(Artist[i].Members)>1{
			for _, v := range Artist[i].Members {
				if strings.Contains(strings.ToLower(v),s){
					_Artist=append(_Artist, Artist[i])
					break 
				}
			}
		}else if loc{
			for _, d := range Location[i].Locations {
				lcc:=strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(d,"_"," "),"-"," "))
				if strings.Contains(lcc,s){
					_Artist=append(_Artist, Artist[i])
					break
				}
			}
		}else if !arrt && !mem && !loc{
			l:=len(_Artist)
			if strings.Contains(strings.ToLower(Artist[i].Name) ,s) || strings.Contains(strconv.Itoa(Artist[i].CreationDate),s) ||strings.Contains(strings.ToLower(Artist[i].FirstAlbum),s ){
				_Artist=append(_Artist, Artist[i])
			}else if len(Artist[i].Members)>1{
				for _, v := range Artist[i].Members {
					lcc:=strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(v,"_"," "),"-"," "))
					if strings.Contains(lcc,s){
						_Artist=append(_Artist, Artist[i])
						break 
					}
				}
			}
			if len(_Artist)==l{
				for _, d := range Location[i].Locations {
					if strings.Contains(strings.ToLower(d),s){
						_Artist=append(_Artist, Artist[i])
						break
					}
				}
			}
		}
		
	}
	return _Artist
}


//**************************************** GET ALL LOCATION******************************
func GetArrayloc(Loca []models.Location,Artists []models.Artist) []string {
	Loc := []string{}
	Item := []string{}
	Members := []string{}
	for _, v := range Loca {
		Loc = append(Loc, v.Locations...)
	}
	Loc=removeDuplicateStr(Loc)
	for _, v := range Artists {
		Item = append(Item, v.Name+" - artist/band")
		Item = append(Item, v.FirstAlbum)
		Members = append(Members, v.Members...)

		for _, t := range Members {
			Item = append(Item, t+" - member")
		}

		Item = append(Item, strconv.Itoa(v.CreationDate))
	}
	Item=removeDuplicateStr(Item)
	for i := 0; i < len(Loc); i++ {
		Loc[i]=strings.ReplaceAll(strings.ReplaceAll(Loc[i],"_"," "),"-"," ")+ " - location"
	}
	Item=append(Item, Loc...)
	return Item
}

// *************************************** FECTH ERROR **********************************
func FecthError(ch <-chan error) bool {
	for err := range ch {
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return false
}

// ************************************** REMOVE THE DUP *********************************
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func FilterLoc(_items []string, s string) ([]string) {
	items:=[]string{}
	for _, v := range _items {
		if strings.Contains(strings.ToLower(v) ,s){
			items = append(items, v)
		}
	}
	return items
}