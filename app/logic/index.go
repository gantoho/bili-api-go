package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to Bili-Favorites API")
}
