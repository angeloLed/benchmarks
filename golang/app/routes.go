package app

import (
	"app/app/controllers"
	// "gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) Init(gr *gin.Engine) {

	userController := controllers.GetInstance()
	// gr.GET("/users", userController.GetAll)
	gr.POST("/users", userController.Store)
}
