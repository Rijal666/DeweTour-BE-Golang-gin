package handlers

import (
	profiledto "dewetour/dto/profile"
	resultdto "dewetour/dto/result"
	"dewetour/models"
	"dewetour/repositories"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func (h *handlerProfile) GetProfile(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(int(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Profile data successfully obtained", Data: convertResponseProfile(profile)})
	
}

func (h *handlerProfile) CreateProfile(c *gin.Context) {
	dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file", dataFile)

	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := profiledto.ProfileRequest{
		ID: int(userId),
		Image: dataFile,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	profile := models.Profile{
		ID: request.ID,
		UserID: int(userId),
		Image: request.Image,
	}

	profile, err = h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	profile, _ = h.ProfileRepository.GetProfile(profile.ID)

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "profile data created successfully", Data: convertResponseProfile(profile)})
	
}

func (h *handlerProfile) DeleteProfile(c *gin.Context)  {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.ProfileRepository.GetProfile(int(userId))
	if err != nil {
		 c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		 return
	}

	fileName := user.Image
	dirPath := "uploads"

	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)

	data, err := h.ProfileRepository.DeleteProfile(user)
	if err != nil {
		 c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		 return
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Failed to delete file"+fileName+":", err)
		 c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		 return
	}

	fmt.Println("File " + fileName + " deleted successfully")

	 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Profile data deleted successfully", Data: convertResponseProfile(data)})
	 
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse  {
	return profiledto.ProfileResponse{
		ID: u.ID,
		UserID: u.UserID,
		User: u.User,
		Image: u.Image,
	}
	
}