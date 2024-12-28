package router

import (
	"CRUD_Gin/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	v1 := r.Group("/v1/api")
	{
		v1.GET("/users", controllers.NewUserController().Index)
		v1.GET("/user/:cccd", controllers.NewUserController().InfoUser)
		v1.POST("/user/add", controllers.NewUserController().AddUser)
	}

	return r
}
