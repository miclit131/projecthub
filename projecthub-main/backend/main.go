package main

import (
	"net/http"
	"projecthub/controller"
	"github.com/gin-gonic/gin"
)

//custom middleware function to allow CORS since our frontend is running on another port / origin.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	server := gin.Default()
	server.Use(CORSMiddleware())
	// Simple api group - dont forget to change the request on postman ;)
	api := server.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
		})
		api.GET("/examples", controller.GetExampleProjects)
		api.POST("/project", controller.CreateProject)
		api.GET("/project/:projectid", controller.GetProjectById)
		api.GET("/projects", controller.GetAllProjects)

	}

	server.Run(":8080")
}
