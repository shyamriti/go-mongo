package controller

import (
	"context"
	"fmt"
	"go-mongo/pkg/database"
	"go-mongo/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during binding data:%v", err).Error())
		return
	}

	name := c.Query("name")
	filter := bson.D{{Key: "name", Value: name}}
	u := bson.D{{Key: "$set", Value: bson.D{{Key: "dob", Value: user.DOB}, {Key: "description", Value: user.Description}, {Key: "address", Value: user.Address}}}}

	_, err := database.Coll.UpdateOne(context.TODO(), filter, u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during updating data:%v", err).Error())
		return
	}
	c.JSON(200, gin.H{
		"dob":         user.DOB,
		"address":     user.Address,
		"description": user.Description,
	})
}
