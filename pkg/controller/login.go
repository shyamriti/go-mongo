package controller

import (
	"context"
	"fmt"
	"go-mongo/pkg/database"
	"go-mongo/pkg/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LogIn(c *gin.Context) {
	var user, result models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("err:%v", err))
		return
	}

	err := database.Coll.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("err:%v", err))
		return
	}

	jwtWrapper := JwtWrapper{
		SecretKey:        "secret",
		Issuer:           "AuthService",
		ExpirationMinute: 1,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	c.SetCookie("user", signedToken, 3600, "/", "localhost", false, true)
	var cookie string
	cookie, err = c.Cookie("user")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fmt.Errorf("err:%v", err))
		return
	}

	c.JSON(http.StatusOK, cookie)
}
