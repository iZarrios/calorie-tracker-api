package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"www.github.com/iZarrios/calorie-tracker-api/pkg/routes"
)

func main() {
	var PORT string
	flag.StringVar(&PORT, "port", "8000", "post of which the server is going to serve on")
	flag.Parse()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())

	r.GET("ping", routes.Ping)
	r.POST("/entry/create", routes.Ping)
	r.GET("/entries", routes.Ping)
	r.GET("/entries/:id/", routes.Ping)
	r.GET("/ingredients/:ingredient/", routes.Ping)

	r.PUT("/entry/update/:id", routes.Ping)
	r.PUT("/ingredients/update/:id", routes.Ping)
	r.DELETE("/entry/delete/:id", routes.Ping)

	r.Use(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})

	r.Run(":" + PORT)
}
