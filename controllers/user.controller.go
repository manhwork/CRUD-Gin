package controllers

import (
	"CRUD_Gin/models"
	"CRUD_Gin/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

// /v1/api/users
func (uc *UserController) Index(c *gin.Context) {
	users, err := services.NewUserService().FindAllUsers()
	if err != nil {
		c.JSON(404, gin.H{"message": "Not found users"})
		return
	}
	c.JSON(200, users)
}

// /v1/api/users/:id
func (uc *UserController) InfoUser(c *gin.Context) {
	cccd := c.Param("cccd")

	user, err := services.NewUserService().FindUserById(cccd)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	c.JSON(200, user)
}

// /v1/api/user/add
func (uc *UserController) AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Invalid input"})
		return
	}

	if err := services.NewUserService().Create(&user); err != nil {
		c.JSON(500, gin.H{"message": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}
