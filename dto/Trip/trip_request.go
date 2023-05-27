package tripdto

type CreateTripRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title" gorm:"type: varchar(255)" validate:"required"`
	CountryID      int    `json:"country_id" gorm:"type: int" validate:"required"`
	Accomodation   string `json:"accomodation" gorm:"type: varchar(255)" validate:"required"`
	Transportation string `json:"transportation" gorm:"type: varchar(255)" validate:"required"`
	Eat            string `json:"eat" gorm:"type: varchar(255)" validate:"required"`
	Day            string `json:"dat" gorm:"type: varchar(255)" validate:"required"`
	Night          string `json:"night" gorm:"type: varchar(255)" validate:"required"`
	DateTrip       string `json:"date_trip" gorm:"type: varchar(255)" validate:"required"`
	Price          string `json:"price" gorm:"type: varchar(255)" validate:"required"`
	Quota          string `json:"quota" gorm:"type: varchar(255)" validate:"required"`
	Description    string `json:"description" gorm:"type: varchar(255)" validate:"required"`
	Photo          string `json:"photo" gorm:"type: varchar(255)" validate:"required"`
}
type UpdateTripRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	CountryID      int    `json:"counter_id"`
	Accomodation   string `json:"accomodation"`
	Transportation string `json:"transportation"`
	Eat            string `json:"eat"`
	Day            string `json:"dat"`
	Night          string `json:"night"`
	DateTrip       string `json:"date_trip"`
	Price          string `json:"price"`
	Quota          string `json:"quota"`
	Description    string `json:"description"`
	Photo          string `json:"photo"`
}