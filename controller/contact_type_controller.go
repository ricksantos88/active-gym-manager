package controller

import (
	"active-gym-manager/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactTypeController struct {
	contactTypeUsecase usecase.ContactTypeUsecase
}

func NewContactTypeController(usercase usecase.ContactTypeUsecase) ContactTypeController {
	return ContactTypeController{
		contactTypeUsecase: usercase,
	}
}

func (ctc *ContactTypeController) GetAllContactTypes(ctx *gin.Context) {
	contactTypes, err := ctc.contactTypeUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, contactTypes)

}
