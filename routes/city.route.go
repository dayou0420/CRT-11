package routes

import (
	"example.com/crt-11/controllers"
	"github.com/gin-gonic/gin"
)

func CityRoute(r *gin.Engine) {
	r.GET("/cities", controllers.GetAllCities())
	r.GET("/city/:cityId", controllers.GetCity())
	r.POST("/city", controllers.CreateCity())
	r.PUT("/city/:cityId", controllers.UpdateCity())
	r.DELETE("/city/:cityId", controllers.DeleteCity())
}
