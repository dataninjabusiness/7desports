package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
	_ "github.com/heroku/x/hmetrics/onload"
)


func ProfileHandler(c *gin.Context) {
  c.Header("Content-Type", "text/html")
  c.HTML(http.StatusOK, "profile.html", nil)
}

func FollowProfileHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "FollowProfile handler not implemented yet",
  })
}

func WallPostHandler(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "post": "WallPost handler not implemented yet",
  })
}

func LikeWallPostHandler(c *gin.Context) {
  c.Header("COntent-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "LikeWallPost handler not implemented yet",
  })
}


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	// router.Static("/views", "views")

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
    api.GET("/profile", ProfileHandler)
    api.POST("/profile/follow/:profileID", FollowProfileHandler)
    api.POST("/profile/wallPost/:profileID", WallPostHandler)
    api.POST("/profile/wallPost/like/:wallPostID", LikeWallPostHandler)
  }

	router.Run(":" + port)
}
