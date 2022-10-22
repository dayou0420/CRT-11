package routes

import (
	"example.com/crt-11/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers())
	router.GET("/user/:userId", controllers.GetUser())
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.EditUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
}
