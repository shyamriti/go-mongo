package controller

import (
	"context"
	"fmt"
	"go-mongo/pkg/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindAllUser(c *gin.Context) {
	var results []map[string]interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch all documents from the collection
	cursor, err := database.Coll.Find(ctx, bson.D{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("err:%v", err))
		return
	}
	for cursor.Next(ctx) {
		var document map[string]interface{}
		err := cursor.Decode(&document)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("err:%v", err))
			return
		}
		results = append(results, document)
	}

	c.JSON(200, results)
}
