package main

import (
	"active-gym-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// dbConn, err := database.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }

	// ContactTypeRepository := repository.NewContactTypeRepository(dbConn)
	// ContactTypeUsecase := usecase.NewContactTypeUsecase(ContactTypeRepository)
	// ContactTypeController := controller.NewContactTypeController(ContactTypeUsecase)

	// server.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// server.GET("/contact-type", ContactTypeController.GetAllContactTypes)
	routes.SetupRouter()
	server.Run(":8001")
}
