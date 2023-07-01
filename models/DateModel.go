package models

type Date struct {
	Id    int
	Dates []string
}
type IndexDate struct {
	Indexes []Date `json:"index"`
}
// //****************************** GET ARTIST METHODE ******************************************************
// func (AllArtists *Artist) GetAllDates() bool {
// 	err := helper.GetJson(Url["date"], &AllArtists)
// 	return err != nil
// }
