package routes

import (
	"active-gym-manager/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	router := gin.Default()

	dbConn, err := database.ConnectDB()
	if err != nil {
		panic(err)

	}

	SetupContactTypeRouter(router, dbConn)

	// return router
}
