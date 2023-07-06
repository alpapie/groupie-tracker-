package models

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
	Artistname		string
}

type Index struct {
	Indexes []Relation `json:"index"`
}


// //****************************** GET ARTIST METHODE ******************************************************
// func (AllRelation *Index) GetAllRelation() bool {
// 	err := helper.GetJson(Url["relation"], &AllRelation)
// 	return err != nil
// }

// //****************************** GET ARTIST METHODE ******************************************************
// func (AllRelation *Index) GetOneRelation() bool {
// 	err := helper.GetJson(Url["relation"], &AllRelation)
// 	return err != nil
// }

