package routers

import (
	"cars-gorm/controllers"

	"github.com/gin-gonic/gin"

	_ "cars-gorm/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @licence.url http://www.apache.org/licences/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:id", controllers.UpdateCar)

	router.GET("/cars/:id", controllers.GetCar)

	router.GET("/cars/allcars", controllers.GetAllCar)

	router.DELETE("/cars/:id", controllers.DeleteCar)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
