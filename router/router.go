package router

import "github.com/gin-gonic/gin"

func Initialize() {

	router * gin.Engine := gin.Default()
	
	initializeRoutes()

	router.Run() // listen and serve on 0.0.0.0:8080

}
