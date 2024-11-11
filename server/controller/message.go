package controller

import (
	"log"
	"message-app/server/database"
	"message-app/server/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)
var collectionMsg = database.DB.Collection("message")
var Validate *validator.Validate
var upgrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request)bool{return true},
}

func ChatWS(c *gin.Context){
	ctx := c.Request.Context()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("gagal konek ke websocket")
		return
	}
	defer conn.Close()

	for{
		var reqMessage models.Message
		err := conn.ReadJSON(&reqMessage)
		if err != nil {
			log.Println("gagal membaca models pada connection websocket")
			break
		}
		reqMessage.Timestamp = time.Now()
		err =	validate.Struct(reqMessage)
		if err !=  nil {
			log.Println("datanya blm bener bray")
			continue
		}
		_, err = collectionMsg.InsertOne(ctx, reqMessage)
		if err != nil {
			log.Println("gagal masukan data ke database", err)
			continue
		}
		err = conn.WriteJSON(reqMessage)
		if err != nil {
			log.Println("error ketika mengirim json ke client", err)
			break
		}
	}
}

