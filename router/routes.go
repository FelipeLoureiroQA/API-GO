package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 *gin.RouterGroup:=router.Group(relativePath:"/api/v1")
	{
		v1.GET(relativePath:"/opening", handlers...:func(ctx *gin.Context)
	ctx.Json(code: http.statusOK, obj:gin.H{
		"msg": "GET Opening",

	})
	)}
}