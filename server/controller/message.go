package controller

import (
	"context"
	"message-app/server/database"
	"message-app/server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
var collectionMsg = database.DB.Collection("message")

func GetMessage(c *gin.Context){
	cur, err := collectionMsg.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mendapatkan data"})
		return
	}

	for cur.Next(context.Background()){
		
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil di akses"})
}

func PostMessage(c *gin.Context){
	var reqMessage models.Message
	err := c.BindJSON(&reqMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "gagal bind request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": reqMessage.Message})
}