package routes

import (
	"active-gym-manager/controller"
	"active-gym-manager/internal/repository"
	"active-gym-manager/internal/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupContactTypeRouter(router  *gin.Engine, dbConn *sql.DB) {

	repository := repository.NewContactTypeRepository(dbConn)
	usecase := usecase.NewContactTypeUsecase(repository)
	contactTypeController := controller.NewContactTypeController(usecase)

	router.GET("/contact-type", contactTypeController.GetAllContactTypes)
}