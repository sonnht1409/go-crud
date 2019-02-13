package controllers

import (
	"../forms"
	"../models"

	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data forms.User
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}
	userRes, userErr := userModel.FindByUsername(data.Username)

	if &userRes != nil || userErr != nil {
		c.JSON(400, gin.H{"message": "username exist"})
		c.Abort()
		return
	}

	err := userModel.Create(data)
	if err != nil {
		c.JSON(400, gin.H{"message": "User could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "User Created"})
}

func (user *UserController) FindByUsername(c *gin.Context) {
	username := c.Param("username")
	userRes, err := userModel.FindByUsername(username)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"id": userRes.ID, "username": userRes.Username})
	}
}

func (user *UserController) SignIn(c *gin.Context) {
	var data forms.User
	if c.BindJSON(&data) != nil {
		c.JSON(400, gin.H{"message": "Invalid form", "form": data})
		c.Abort()
		return
	}

	userRes, err := userModel.FindByUsername(data.Username)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found", "error": err.Error()})
		c.Abort()
		return
	}

	if data.Password != userRes.Password || &userRes == nil {
		c.JSON(401, gin.H{"message": "Unauthorized, wrong username or password"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"id": userRes.ID, "username": userRes.Username})
}
