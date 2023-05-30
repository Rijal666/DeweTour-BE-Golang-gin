package handlers

import (
	resultdto "dewetour/dto/result"
	transactiondto "dewetour/dto/transaction"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type HandleTransaction struct {
	TransactionRepository repositories.TransactionRepository
	TripRepository repositories.TripRepository
	userRepository repositories.UserRepository
}

func NewHandleTransaction(TransactionRepository repositories.TransactionRepository, TripRepository repositories.TripRepository, UserRepository repositories.UserRepository) *HandleTransaction {
	return &HandleTransaction{ TransactionRepository,TripRepository, UserRepository}
}

func (h *HandleTransaction) GetTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	response := resultdto.SuccessResult{Status: http.StatusOK, Data: transaction}
	c.JSON(http.StatusOK, response)
}

func (h *HandleTransaction) CreateTransaction(c *gin.Context) {
	// var request transactiondto.TransRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	// 	return
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	dataFile := c.GetString("dataFile")

	counterqty, _ := strconv.Atoi(c.PostForm("counter_qty"))
	total, _ := strconv.Atoi(c.PostForm("total"))
	tripid, _ := strconv.Atoi(c.PostForm("trip_id"))

	request := transactiondto.TransactionResponse{
		Name: c.PostForm("name"),
		Gender: c.PostForm("gender"),
		Phone: c.PostForm("phone"),
		CounterQty: counterqty,
		Total:      total,
		Status:     c.PostForm("status"),
		Attachment: dataFile,
		TripID:     tripid,
		UserID:     int(userId),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tripID, err := h.TripRepository.GetTrip(request.TripID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		fmt.Println("error 1")
	}	
	user, _ := h.userRepository.GetUser(request.UserID)

	transaction := models.Transaction{
		Name: request.Name,
		Gender: request.Gender,
		Phone: request.Phone,
		CounterQty: request.CounterQty,
		Total:      request.Total,
		Status:     request.Status,
		TripId:     request.TripID,
		Trip: CovertTripResponse(tripID),
		UserID:     request.UserID,
		User: ConvertResponseUser(user),
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := resultdto.SuccessResult{Status: http.StatusOK, Data: data}
	c.JSON(http.StatusOK, response)

}

// validation := validator.New()
// if err := validation.Struct(request); err != nil {
// 	c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	return
// }

func (h *HandleTransaction) FindTransaction(c *gin.Context) {
	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: transaction})
}

func (h *HandleTransaction) DeleteTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Data: data})
}

func CovertTripResponse(c models.Trip) models.TripResponse{
	return models.TripResponse{
		ID: c.ID,
		Title: c.Title,
		CountryID: c.CountryId,
		Country: c.Country,
		Accomodation: c.Accomodation,
		Transportation: c.Transportation,
		Eat: c.Eat,
		Day: c.Day,
		Night: c.Night,
		DateTrip: c.DateTrip,
		Price: c.Price,
		Quota: c.Quota,
		Description: c.Description,
		Image: c.Image,
	}
}

func ConvertResponseUser(u models.User) models.UsersProfileResponse  {
	return models.UsersProfileResponse{
		ID: u.ID,
		Fullname: u.Fullname,
		Email: u.Email,
		Phone: u.Phone,
		Address: u.Address,
	}
	
}