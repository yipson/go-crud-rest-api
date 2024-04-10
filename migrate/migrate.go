package main

import (
	"github.com/yipson/go-crud/initializers"
	"github.com/yipson/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
