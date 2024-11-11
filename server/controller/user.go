package controller

import (
	"context"
	"message-app/server/database"
	"message-app/server/middleware"
	"message-app/server/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var collectionUsr = database.DB.Collection("user")
var ReqErr = gin.H{"message":"data yang dimasukan tidak sesuai format"}

// register
func Register(c *gin.Context){
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ReqErr)
		return
	}

	count, err := collectionUsr.CountDocuments(context.Background(), bson.M{"username": newUser.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mencari username"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "username sudah ada"})
		return
	}

	_, err = collectionUsr.InsertOne(context.Background(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal upload data"})
		return
	}
	c.IndentedJSON(http.StatusOK, newUser)
}

// login
func Login(c *gin.Context){
	var checkUser models.User
	err := c.BindJSON(&checkUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ReqErr)
		return
	}

	var user models.User
	err = collectionUsr.FindOne(context.Background(), bson.M{
		"username": checkUser.Username, 
		"password": checkUser.Password,
	}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "username atau password salah"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mendapatkan user"})
		return
	}
	token, err := middleware.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal generate token"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"user": user, "token": token, "message": "berhasil login"})
}