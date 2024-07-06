package controllers

import (
	"fmt"
	"myapp/initializers"
	"myapp/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var db = initializers.ConnectToDB()
	//Get Data
	//Create a post
	//return

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
	}

	userIDUint, err := strconv.ParseUint(fmt.Sprintf("%v", userId), 10, 32)
	if err != nil {
		panic(err)
	}

	newPost := models.Post{Title: body.Title, Body: body.Body, UserID: uint(userIDUint)}

	result := db.Create(&newPost)

	if result.Error != nil {
		// c.Status(404)
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to connect",
			"data":    gin.H{},
		})

	}

	c.JSON(200, gin.H{
		"post": newPost,
	})
}

func PostsIndex(c *gin.Context) {

	var db = initializers.ConnectToDB()

	//Get the post

	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
	}
	var user models.User
	db.Preload("Posts").First(&user, userId)

	fmt.Println("user Id", userId)
	c.JSON(200, gin.H{
		"posts": user.Posts,
	})

}

func PostsShow(c *gin.Context) {
	var db = initializers.ConnectToDB()
	//Get id from url
	id := c.Param("id")
	//Get the post

	//respond with them

	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(401, gin.H{
			"success": false,
			"message": "Invalid user id",
		})
	}

	userIDUint, err := strconv.ParseUint(fmt.Sprintf("%v", userId), 10, 32)
	if err != nil {
		panic(err)
	}

	var post models.Post
	var user models.User
	db.First(&post)

	postId := []string{id}

	db.Preload("Posts", "id IN (?)", postId).First(&user, userIDUint)
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	var db = initializers.ConnectToDB()
	//Get id from url
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//Get the post

	//respond with them
	var post models.Post
	db.First(&post, id)

	db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	var db = initializers.ConnectToDB()
	//Get id from url
	id := c.Param("id")

	db.Delete(&models.Post{}, id)

	c.JSON(400, gin.H{
		"success": true,
		"message": "Post deleted",
		"data":    gin.H{},
	})
}
