package models

import "time"

type Transaction struct {
	ID         int            `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int            `json:"counter_Qty" gorm:"type int"`
	Total      int            `json:"total" gorm:"type int"`
	Status     string         `json:"status" gorm:"type varchar(255)"`
	Attachent  string         `json:"attachment" gorm:"type varchar(255)"`
	Trip       []TripResponse `json:"trip" gorm:"many2many:TripID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
}

type TransactionResponse struct {
	ID int 
	CounterQty int            `json:"counter_Qty"`
	Total      int            `json:"total"`
	Status     string         `json:"status"`
	Attachent  string         `json:"attachment"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}