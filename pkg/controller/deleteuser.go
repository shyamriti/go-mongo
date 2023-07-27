package controller

import (
	"context"
	"fmt"
	"go-mongo/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c *gin.Context) {

	name := c.Query("name")
	filter := bson.D{{Key: "name", Value: name}}
	resp, err := database.Coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during deleting data:%v", err).Error())
		return
	}
	c.JSON(200, resp)
}
