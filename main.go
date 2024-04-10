package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yipson/go-crud/controllers"
	"github.com/yipson/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}
