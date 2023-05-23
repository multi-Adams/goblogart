package main

import (
	"goblogart/controllers"
	"goblogart/inits"
	"goblogart/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {

	r := gin.Default()

	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)

	r.POST("/users", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/users", controllers.GetUsers)
	r.GET("/auth", middlewares.RequireAuth, controllers.Validate)

	r.Run()
}
