package router

import (
	"go-mongo/pkg/controller"
	"go-mongo/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	users := r.Group("/user")

	// Create a new user
	users.POST("/register", controller.AddUser)

	// user login
	users.POST("/login", controller.LogIn)

	user := users.Use(middleware.IsAuthorized)

	// user logout
	user.POST("/logout", controller.LogOut)

	// Get user data
	user.GET("/find", controller.FindUser)
	user.GET("/findall", controller.FindAllUser)

	// Update user data
	user.PUT("/update", controller.UpdateUser)

	// Delete user data
	user.DELETE("/delete", controller.DeleteUser)

	return r

}
