package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerUser(userRepository, profileRepository)

	r.GET("/users", middleware.Auth(h.FindUsers))
	r.GET("/user", middleware.Auth(h.GetUser))
	r.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}