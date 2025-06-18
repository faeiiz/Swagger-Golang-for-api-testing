package main

import (
	"swag/controllers"
	_ "swag/docs"
	"swag/initializers"
	"swag/migrate"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server for a Gin application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	migrate.Migrate()
}
func main() {
	r := gin.Default()
	// r.Static("/swaggerui", "./swaggerui")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(sw))
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.Run()
}
