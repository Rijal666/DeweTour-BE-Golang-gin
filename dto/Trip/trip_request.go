package tripdto

type CreateTripRequest struct {
	ID             int    `json:"id"`
	Title          string `json:"title" form:"title" validate:"required"`
	CountryId      int    `json:"country_id" validate:"required" form:"country_id"`	
	Accomodation   string `json:"accomodation" form:"accomodation" validate:"required"`
	Transportation string `json:"transportation" form:"transportation" validate:"required"`
	Eat            string `json:"eat" form:"eat" validate:"required"`
	Day            int    `json:"day" form:"day" validate:"required"`
	Night          int    `json:"night" form:"night" validate:"required"`
	DateTrip       string `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int    `json:"price" form:"price" validate:"required"`
	Quota          int    `json:"quota" form:"quota" validate:"required"`
	Description    string `json:"description" form:"description" validate:"required"`
	Image          string `json:"image" form:"image" validate:"required"`
}

type Update1TripRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title" form:"title"`
	CountryId int    `json:"country_id" form:"country_id"`
	// Country        models.CountryResponse `json:"country"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transportation" form:"transportation"`
	Eat            string `json:"eat" form:"eat"`
	Day            int    `json:"day" form:"day"`
	Night          int `json:"night" form:"night"`
	DateTrip       string `json:"date_trip" form:"date_trip"`
	Price          int    `json:"price" form:"price"`
	Quota          int    `json:"quota" form:"quota"`
	Description    string `json:"description" form:"description"`
	Image          string `json:"image" form:"image"`
}