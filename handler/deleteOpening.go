package handler

import (
	"fmt"
	"net/http"
	"github.com/FelipeLoureiroQA/API-GO/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
			"queryParameter").Error())
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
}

    if err := db.Delete(&opening).Error; err!= nil {
	sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("opening with id %s not found", id))
        return
    }

	sendSuccess(ctx, "delete-opening", opening)
}