package controllers

import (
	"net/http"

	"app/app/services"
	// "gopkg.in/gin-gonic/gin.v1"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	// "storageGo/app/transformers"
)

type UserController struct {
	Service services.UserService
}

func GetInstance() UserController {
	service := services.UserService{}
	service.Init()
	return UserController{Service: service}
}

func (u *UserController) GetAll(c *gin.Context) {
	users := u.Service.ShowMany()
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func(u *UserController) Store(c *gin.Context) {
	
	body := bson.M{}
	c.BindJSON(&body)

	user, err := u.Service.Store(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":  user,
	   })
}
