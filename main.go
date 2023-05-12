package main

import (
	"github.com/FadhliAbrar/belajar-golang-jwt/controller"
	"github.com/FadhliAbrar/belajar-golang-jwt/database"
	"github.com/FadhliAbrar/belajar-golang-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:kraken1288@tcp(localhost:3306)/belajar_golang_jwt?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
