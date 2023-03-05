package controllers

import (
	"attendance-app/initializers"
	"attendance-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data from body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create post obj
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	//get data
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsGetAll(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostGet(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	err := initializers.DB.First(&post, id).Error //fetch error kalo not found

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"post": post,
		})
	}
}
