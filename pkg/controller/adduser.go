package controller

import (
	"context"
	"fmt"
	"go-mongo/pkg/database"
	"go-mongo/pkg/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *gin.Context) {
	var user models.User
	user.CreatedAt = time.Now().Format("01-02-2006 15:04:05")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during binding:%v", err))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}
	user.Password = string(hashedPassword)

	result, err1 := database.Coll.InsertOne(context.TODO(), user)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("error during inserting data:%v", err1).Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": result.InsertedID,
	})
}
