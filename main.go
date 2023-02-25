package main

// "net/http"
// "github.com/gin-gonic/gin"

import (
	"attendance-app/controllers"
	"attendance-app/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)
	r.PUT("/update/:id", controllers.PostUpdate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
