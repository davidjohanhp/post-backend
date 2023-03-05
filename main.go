package main

import (
	"attendance-app/controllers"
	"attendance-app/initializers"
	"attendance-app/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&models.Post{})

}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)
	r.PUT("/update/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.GET("/get-post/:id", controllers.PostGet)
	r.GET("/get-posts", controllers.PostsGetAll)
	r.Run()
}
