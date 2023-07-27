package controller

import (
	"bytes"
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

func TestUpdateUser(t *testing.T) {
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
	router.PUT("/user/update", UpdateUser)

	// Define the user ID to update
	userName := "john"

	// Define the updated user details
	updatedUser := models.User{
		DOB:         "87368612",
		Address:     "agartala",
		Description: "jsdggh",
	}
	requestBody, _ := json.Marshal(updatedUser)

	// Create a PUT request to update the user
	req, _ := http.NewRequest("PUT", "/user/update?name="+userName, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request and record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Assert the response body contains the updated user details (using your mock data)

}
