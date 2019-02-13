package main

import (
	"net/http"

	"./controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		user := new(controllers.UserController)
		v1.POST("/users/sign-up", user.Create)
		v1.POST("users/sign-in", user.SignIn)
		v1.GET("/users/:username", user.FindByUsername)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	router.Run()
}
