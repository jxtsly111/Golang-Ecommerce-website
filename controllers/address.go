package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jxtsly111/ecommerce-yt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error":"Invalid Code"})
			c.Abort()
			return
		}
		address, err := ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "Intrrnal server error")
		}

		var addresses models.Address
		addresses.Address_id = primitive.NewObjectID()
		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value:bson.D{primitive.E{Key: "path",Value: "$address"}}}}
		group := bson.D{{Key: "$group",Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"},{Key: "count", Value: bson.D{primitive.E{Key: "$sum, Value:1 "}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter,unwind,group})
		if err != nil {
			c.IndentedJSON(500, "Internal server error ")
		}

	}
}

func EditHomeAddress() gin.HandlerFunc {
	
}

func EditWorkAddress() gin.HandlerFunc {
	
}

func DeleteAddress() gin.HandlerFunc {
	return func (c *gin.Context)  {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
	}

	addresses := make([]models.Address,0)
	usert_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.IndentedJSON(500, "Internal Server Error",)
	}

	var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
	update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
	_, err = UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.IndentedJSON(404, "Wrong Command")
		return
	}
	defer cancel()
	ctx.Done()
	c.IndentedJSON(200, "Successfully Deleted")
}
}