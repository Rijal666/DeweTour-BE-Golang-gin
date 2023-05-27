package main

import (
	routes "dewetour/Routes"
	"dewetour/migration"
	"dewetour/pkg/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	mysql.AutoMigrate()
	migration.RunAutoMigrate()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.RouteInit(r.Group("/api/v1"))

	// r.Static("/uploads", "./uploads")

	fmt.Println("Server Started")
	http.ListenAndServe("localhost:5000", r)
}
