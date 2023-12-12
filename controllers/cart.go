package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

	type Application struct {
		prodCollection *mongo.Collection
		userCollection *mongo.Collection
	}

	func NewApplication(prodCollection, userCollection *mongo.Collection) *Application{
		return &Application{
			prodCollection: prodCollection,
			userCollection: userCollection,
		}
	}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID , err := primitive.ObjectIDFromHex(productQueryID)

		if err!= nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

func RemoteItem() gin.HandlerFunc {
	
}

func GetItemFromCart() gin.HandlerFunc {
	
}

func BuyFromCart() gin.HandlerFunc{

}

func InstantBuy() gin.HandlerFunc{

}