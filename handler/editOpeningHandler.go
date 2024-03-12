package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})

}
