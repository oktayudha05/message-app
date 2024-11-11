package controller

import (
	"message-app/server/database"
	"message-app/server/middleware"
	"message-app/server/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var collectionUsr = database.DB.Collection("user")
var ReqErr = gin.H{"message":"data yang dimasukan tidak sesuai format"}
var validate *validator.Validate

func init(){
	validate = validator.New()
}

// register
func Register(c *gin.Context){
	ctx := c.Request.Context()
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ReqErr)
		return
	}
	err = validate.Struct(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data tidak lengkap"})
		return
	}

	count, err := collectionUsr.CountDocuments(ctx, bson.M{"username": newUser.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mencari username"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "username sudah ada"})
		return
	}

	_, err = collectionUsr.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal upload data"})
		return
	}
	c.IndentedJSON(http.StatusOK, newUser)
}

// login
func Login(c *gin.Context){
	ctx := c.Request.Context()
	var checkUser models.User
	err := c.BindJSON(&checkUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ReqErr)
		return
	}
	err = validate.Struct(checkUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "data tidak lengkap"})
		return
	}

	var user models.User
	err = collectionUsr.FindOne(ctx, bson.M{
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