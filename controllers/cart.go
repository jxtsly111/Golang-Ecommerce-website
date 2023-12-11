package controllers

import (
	"github.com/gin-gonic/gin"
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

func AddToCart() gin.HandlerFunc {
	
}

func RemoteItem() gin.HandlerFunc {
	
}

func GetItemFromCart() gin.HandlerFunc {
	
}

func BuyFromCart() gin.HandlerFunc{

}

func InstantBuy() gin.HandlerFunc{

}