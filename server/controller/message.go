package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context){
	// collection := database.DB.Collection("message")


	// cur, err := collection.Find(context.Background(), bson.M{})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mendapatkan data"})
	// 	return
	// }

	// for cur.Next(context.Background()){
		
	// }

	c.JSON(http.StatusOK, gin.H{"message": "berhasil di akses"})
}