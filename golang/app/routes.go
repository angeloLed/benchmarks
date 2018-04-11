package app

import (
	"app/app/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) Init(gr *gin.Engine) {

	userController := controllers.GetInstance()
	gr.GET("/users", userController.GetAll)
	gr.GET("/userHeats", userController.GetAllUserHasHeatZone)
	gr.POST("/users", userController.Store)
}
