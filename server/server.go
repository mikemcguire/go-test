package server

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

const PORT = "8080"


func Server() {

	router := gin.Default()
	registerRoutes(router)
	router.Run(":"+PORT)

	return
}


func registerRoutes(router gin.IRouter){
	router.GET("/search", func(context *gin.Context){
		context.JSON(
			http.StatusOK,
			gin.H{"Your request": "is valid"})
	})
}

