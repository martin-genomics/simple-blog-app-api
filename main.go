package main

import (
	"myapp/controllers"
	"myapp/initializers"
	"myapp/middlware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	// initializers.

}

func main() {
	r := gin.Default()

	// r.RouterGroup.GET()
	v1 := r.Group("/api/v1/user")

	v1.POST("signup", controllers.Signup)
	v1.POST("login", controllers.Login)

	//posts routes
	v1.POST("/posts", middlware.RequireAuth, controllers.PostsCreate)
	v1.GET("/posts", middlware.RequireAuth, controllers.PostsIndex)
	v1.GET("/posts/:id", middlware.RequireAuth, controllers.PostsShow)
	v1.PUT("/posts/:id", middlware.RequireAuth, controllers.PostsUpdate)
	v1.DELETE("/posts/:id", middlware.RequireAuth, controllers.PostsDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
