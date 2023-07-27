package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", false, true)

	cookie, err := c.Cookie("user")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "log out abort")
		return
	}
	if cookie != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "cookies not nil")
		return

	}
	c.JSON(200, "log out successfully")
}
