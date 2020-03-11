package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		} else {
			port = os.Getenv("PORT")
		}
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
