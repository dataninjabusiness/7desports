package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
  
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

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
