package models

//CityList provides the struct for the list
type Cities struct {
	Cities []City `json:"citylist"`
}

//City provides struct for city location information
type City struct {
	ID    int `json:"id"`
	Coord `json:"coord"`
}
