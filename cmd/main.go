package main

import (
	"active-gym-manager/controller"
	"active-gym-manager/infrastructure/database"
	"active-gym-manager/internal/repository"
	"active-gym-manager/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConn, err := database.ConnectDB()
	if err != nil {
		panic(err)

	}

	ContactTypeRepository := repository.NewContactTypeRepository(dbConn)
	ContactTypeUsecase := usecase.NewContactTypeUsecase(ContactTypeRepository)
	ContactTypeController := controller.NewContactTypeController(ContactTypeUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/contact-type", ContactTypeController.GetAllContactTypes)

	server.Run(":8001")
}
