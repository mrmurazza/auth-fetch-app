package main

import (
	userImpl "authapp/domain/user/impl"
	"authapp/handler"
	"authapp/pkg/database"

	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDatabase()

	// init service & repo
	itemRepo := userImpl.NewRepo(database.DB)
	itemSvc := userImpl.NewService(itemRepo)

	// init handler
	apiHandler := handler.NewApiHandler(itemSvc)

	v1 := r.Group("/api/v1", Logger())
	{
		v1.POST("/login", apiHandler.Login)
		v1.POST("/user", apiHandler.CreateUser)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
