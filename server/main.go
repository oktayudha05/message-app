package main

import (
	"message-app/server/controller"
	"message-app/server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "endpoint hitted"})
	})
	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/message", middleware.JWTMiddleware(), controller.GetMessage)
	router.POST("/message", middleware.JWTMiddleware(), controller.PostMessage)

	router.Run(":3000")
}