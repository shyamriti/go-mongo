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

func FindUser(c *gin.Context) {
	var user models.User
	name := c.Query("name")
	filter := bson.D{{Key: "name", Value: name}}
	err := database.Coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during finding data:%v", err).Error())
		return
	}
	c.JSON(200, gin.H{
		"name":        user.Name,
		"dob":         user.DOB,
		"address":     user.Address,
		"description": user.Description,
		"created_at":  user.CreatedAt,
	})
}
