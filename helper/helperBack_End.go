package helper

import (
	"encoding/json"
	"fmt"
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
func GetJson( url string, model interface{}) error {
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

// *************************************** FECTH ERROR **********************************
func FecthError( ch <-chan error) (bool){
	for err := range ch{
		if err!=nil{
			fmt.Println(err)
			return true
		}
	}
	return false
}