package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func ProfileRoutes(r *gin.RouterGroup) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	r.GET("/profile", middleware.Auth(h.GetProfile))
	r.POST("/profile", middleware.Auth(middleware.UploadFile(h.CreateProfile)))
	r.DELETE("/profile", middleware.Auth(h.DeleteProfile))

}