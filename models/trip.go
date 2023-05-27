package models

type Trip struct {
	ID             int             `json:"id" gorm:"primary_key:auto_increment"`
	Title          string          `json:"title" gorm:"type: varchar(255)"`
	CountryID      int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation" gorm:"type: varchar(255)"`
	Transportation string          `json:"transportation" gorm:"type: varchar(255)"`
	Eat            string          `json:"eat" gorm:"type: varchar(255)"`
	Day            string          `json:"day" gorm:"type: varchar(255)"`
	Night          string          `json:"night" gorm:"type: varchar(255)"`
	DateTrip       string          `json:"datetrip" gorm:"type: varchar(255)"`
	Price          string          `json:"price" gorm:"type: varchar(255)"`
	Quota          string          `json:"quota" gorm:"type: varchar(255)"`
	Description    string          `json:"description" gorm:"type: varchar(255)"`
	Photo          string          `json:"photo" gorm:"type: varchar(255)"`
}

type TripResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	CountryID      int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            string          `json:"day"`
	Night          string          `json:"night"`
	DateTrip       string          `json:"datetrip"`
	Price          string          `json:"price"`
	Quota          string          `json:"quota"`
	Description    string          `json:"description"`
	Photo          string          `json:"photo"`
}

func (TripResponse) TableName() string {
	return "trips"
}