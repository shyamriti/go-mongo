package controller

import (
	"context"
	"go-mongo/pkg/database"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
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
	database.Client.Database("userData").Collection("user_test")
	// Set up routes
	router.DELETE("/user/delete", DeleteUser)

	// Define the user ID to delete
	userName := "john"

	// Create a DELETE request to delete the user
	req, _ := http.NewRequest("DELETE", "/user/delete?name="+userName, nil)

	// Perform the request and record the response
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the status code
	assert.Equal(t, http.StatusOK, resp.Code)
}
