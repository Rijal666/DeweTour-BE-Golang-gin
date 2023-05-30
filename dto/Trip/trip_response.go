package tripdto

import "dewetour/models"

type TripResponse struct {
	ID        int `json:"id"`
	Title string `json:"title"`
	CountryID int `json:"country_id"`
	Country   models.CountryResponse `json:"country"`
	Accomodation   string `json:"accomodation" `
	Transportation string `json:"transportation" `
	Eat            string `json:"eat" `
	Day            int `json:"dat" `
	Night          int `json:"night" `
	DateTrip       string `json:"date_trip" `
	Price          int `json:"price" `
	Quota          int `json:"quota" `
	Description    string `json:"description" `
	Image          string `json:"image"`
}