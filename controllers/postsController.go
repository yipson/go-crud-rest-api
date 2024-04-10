package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yipson/go-crud/initializers"
	"github.com/yipson/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	//Get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"message": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get ID of URL
	id := c.Param("id")

	//Get post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Updating
	initializers.DB.Model(&post).Updates(models.Post{
		Body:  body.Body,
		Title: body.Title,
	})

	//Respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get id from URL
	id := c.Param("id")

	//Delete the prost
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
