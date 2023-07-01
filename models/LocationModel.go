package models

type Location struct {
	Id        int
	Locations []string
	Dates     string
}

type IndexLoc struct {
	Indexes []Location `json:"index"`
}

type LocationOne struct {
	LocationOne Location
	Date        Date
}

// //****************************** GET ARTIST METHODE ******************************************************
// func (AllLocation *Location) GetAllRLocation() bool {
// 	err := helper.GetJson(Url["location"], &AllLocation)
// 	return err != nil
// }
