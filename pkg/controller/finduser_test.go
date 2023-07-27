package controller

import (
	"context"
	"encoding/json"
	"go-mongo/pkg/database"
	"go-mongo/pkg/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Set up Gin router
	router := gin.Default()
	err := godotenv.Load("/home/shyam/go-workspace/go-mongo/.env")
	if err != nil {
		log.Print(err.Error())
	}
	dbColl := os.Getenv("TEST_DB_COLLECTION")

	// Create a mock MongoDB client and inject it into the handler
	database.Connection(dbColl)
	defer database.Client.Disconnect(context.Background())

	// Set up routes
	router.GET("/user/find", FindUser)

	// Define the user ID to fetch
	userName := "john"

	// Create a GET request to get a user by ID
	req, _ := http.NewRequest("GET", "/user/find?name="+userName, nil)

	// Perform the request and record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.Code)
	// Assert the response body contains the correct user details (using your mock data)
	var user models.User
	err = json.Unmarshal(resp.Body.Bytes(), &user)
	assert.NoError(t, err)
}
