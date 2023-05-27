package models

type Profile struct {
	ID     int                  `json:"id" gorm:"primary_key:auto_increment"`
	UserID int                  `json:"user_id" gorm:"type:int"`
	User   UsersProfileResponse `json:"user"`
	Image  string               `json:"image" gorm:"type: varchar(255)"`
}

type ProfileResponse struct {
	Image  string `json:"image"`
	UserID string `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}