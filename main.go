package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Static("/views", "views")

	router.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.tmpl.html", nil)
		c.HTML(http.StatusOK, "index.html", nil)
	})

  api := router.Group("/api")
  {
    api.GET("", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "status": "success",
        "message": "pong!",
      })
    })
    // api.GET("/users/all", users.GetAllUsers(db))
    // api.GET("/users/username", users.GetUserByUserName(db))
    // api.POST("/users/create", users.CreateUser(db))
  }

	router.Run(":" + port)
}
