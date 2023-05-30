package transactiondto

type createTransactionRequest struct {
	ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	Name       string `json:"name" validate:"required" form:"name"`
	Gender     string `json:"gender" validate:"required" form:"gender"`
	Phone      string `json:"phone" validate:"required" form:"phone"`
	CounterQty int    `json:"counter_qty" validate:"required" form:"counter_qty"`
	Total      int    `json:"total" validate:"required" form:"total"`
	Status     string `json:"status" validate:"required" form:"status"`
	Attachment string `json:"attachment" validate:"required" form:"attachment"`
	TripID     int    `json:"trip_id" validate:"required" form:"trip_id"`
	UserID     int    `json:"user_id" validate:"required" form:" user_id"`
}

type UpdateTransactionRequest struct {
	ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	Name       string `json:"name" form:"name"`
	Gender     string `json:"gender" form:"gender"`
	Phone      string `json:"phone" form:"phone"`
	CounterQty int    `json:"counter_qty" form:"counter_qty"`
	Total      int    `json:"total" form:"total"`
	Status     string `json:"status" form:"status"`
	Attachment string `json:"attachment" form:"attachment"`
	TripID     int    `json:"trip_id" form:"trip_id"`
	UserID     int    `json:"user_id" form:"user_id"`
}