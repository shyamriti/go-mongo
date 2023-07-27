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

func TestCreateUser(t *testing.T) {
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
	router.POST("/user/add", AddUser)

	// Define the request payload
	newUser := models.User{
		Name:        "john",
		DOB:         "1/1/2000",
		Address:     "agartala",
		Description: "i am a developer",
	}
	requestBody, _ := json.Marshal(newUser)

	// Create a POST request to create a new user
	req, _ := http.NewRequest("POST", "/user/add", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request and record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Assert the response body contains the new user ID
	var response map[string]string
	_ = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NotEmpty(t, response["id"])
}
