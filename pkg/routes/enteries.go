package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/iZarrios/calorie-tracker-api/pkg/models"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

}

func AddEntry(c *gin.Context) models.Entry {
	return models.Entry{}
}
