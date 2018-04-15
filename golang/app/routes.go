package app

import (
	"app/app/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) Init(gr *gin.Engine) {

	heatController := controllers.GetInstance()
	gr.GET("/heats", heatController.GetAll)
	gr.GET("/userHeats", heatController.GetAllUserHasHeatZone)
	gr.POST("/heats", heatController.Store)
}
