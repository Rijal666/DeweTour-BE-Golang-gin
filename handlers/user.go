package handlers

import (
	resultdto "dewetour/dto/result"
	userdto "dewetour/dto/user"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	UserRepository repositories.UserRepository
	ProfileRepository     repositories.ProfileRepository
}

func HandlerUser(UserRepository repositories.UserRepository, ProfileRepository repositories.ProfileRepository) *handler {
	return &handler{
		UserRepository:        UserRepository,
		ProfileRepository:     ProfileRepository,
		// CartRepository:        CartRepository,
		// TransactionRepository: TransactionRepository,
	}
}

func (h *handler) FindUsers(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	if userAdmin {
		users, err := h.UserRepository.FindUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})

		}

		if len(users) > 0 {
			c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Data for all users was successfully obtained", Data: users})
		} else {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "No record found"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: http.StatusUnauthorized, Message: "Sorry, you're not Admin"})
	}
}

func (h *handler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		 c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "User data successfully obtained", Data: user})
}


func (h *handler) DeleteUser(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId, _ := strconv.Atoi(c.Param("id"))
	userAdmin := userLogin.(jwt.MapClaims)["is_admin"].(bool)
	profiles, err := h.ProfileRepository.FindProfiles()

	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}


	if userAdmin {

		for _, profile := range profiles {
			if profile.UserID == int(userId) {
				userProfile, err := h.ProfileRepository.GetProfile(profile.ID)
				if err != nil {
					c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
				}
				fileName := userProfile.Image
				dirPath := "uploads"
	
				filePath := fmt.Sprintf("%s/%s", dirPath,fileName)
	
				_, err = h.ProfileRepository.DeleteProfile(userProfile)
				if err != nil {
					c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest,Message: err.Error()})
				}
				err = os.Remove(filePath)
				if err != nil {
					fmt.Println("Failed to delete file"+fileName+":", err)
					c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError,Message: err.Error()})
				}
	
				fmt.Println("File "+ fileName + " deleted successfully")
			}
		}
	
		user, err := h.UserRepository.GetUser(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		}
		data, err := h.UserRepository.DeleteUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		}
		c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "User data deleted successfully", Data: convertResponse(data)})
	} else {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: "Sorry, you're not Admin"})
	}
}

func  convertResponse(u models.User)  userdto.UserResponse{
	return userdto.UserResponse{
		ID: u.ID,
		Fullname: u.Fullname,
		Email: u.Email,
		Phone: u.Phone,
		Address: u.Address,
	}
	
}