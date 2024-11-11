package controller

import (
	"message-app/server/database"
	"message-app/server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)
var collectionMsg = database.DB.Collection("message")

func GetMessage(c *gin.Context){
	ctx := c.Request.Context()
	cur, err := collectionMsg.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mendapatkan data"})
		return
	}

	for cur.Next(ctx){
		
	}

	c.JSON(http.StatusOK, gin.H{"message": "berhasil di akses"})
}

func PostMessage(c *gin.Context){
	ctx := c.Request.Context()
	var reqMessage models.Message
	err := c.BindJSON(&reqMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "gagal bind request"})
		return
	}
	reqMessage.Timestamp = time.Now()
	err = validate.Struct(reqMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data tidak lengkap"})
		return
	}

	_, err = collectionMsg.InsertOne(ctx, reqMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal kirim pesan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": reqMessage})
}