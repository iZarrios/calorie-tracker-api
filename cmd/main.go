package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"www.github.com/iZarrios/calorie-tracker-api/pkg/routes"
)

func main() {
	// Added it to the .env file instead
	//	var PORT string
	//	flag.StringVar(&PORT, "port", "8000", "post of which the server is going to serve on")

	//	flag.Parse()

	PORT := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())

	r.GET("/ping", routes.Ping)
	r.GET("/entries", routes.GetEntries)
	r.GET("/entries/:id/", routes.GetEntryByID)
	r.GET("/ingredients/:ingredient/", routes.GetEntryByIngredient)

	r.POST("/entry/create", routes.AddEntry)

	r.PUT("/entry/update/:id", routes.UpdateEntry)
	r.PUT("/ingredients/update/:id", routes.UpdateIngredient)
	r.DELETE("/entry/delete/:id", routes.DeleteEntry)

	r.Use(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "route not found",
		})
	})
	r.Run(":" + PORT)
}
