package tripdto

import "dewetour/models"

type TripResponse struct {
	ID        int
	CountryID int
	Country   models.CountryResponse
	Accomodation   string `json:"accomodation" gorm:"type: varchar(255)" `
	Transportation string `json:"transportation" gorm:"type: varchar(255)" `
	Eat            string `json:"eat" gorm:"type: varchar(255)" `
	Day            string `json:"dat" gorm:"type: varchar(255)" `
	Night          string `json:"night" gorm:"type: varchar(255)" `
	DateTrip       string `json:"date_trip" gorm:"type: varchar(255)" `
	Price          string `json:"price" gorm:"type: varchar(255)" `
	Quota          string `json:"quota" gorm:"type: varchar(255)" `
	Description    string `json:"description" gorm:"type: varchar(255)" `
	Photo          string `json:"photo" gorm:"type: varchar(255)"`
}