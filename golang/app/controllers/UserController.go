package controllers

import (
	"net/http"

	"app/app/services"
	"strconv"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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

	var filters bson.M //implementation of the filters in another spacetime

	users := u.Service.ShowMany(filters)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (u *UserController) GetAllUserHasHeatZone(c *gin.Context) {
	var err error
	var x ,y, rad int

	x, err = strconv.Atoi(c.Query("x"))
	y, err = strconv.Atoi(c.Query("y"))
	rad, err = strconv.Atoi(c.Query("radius"))

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "missig x,y,radius",
		})
		return
	}

	users := u.Service.GetAllUserHasHeatZone(x, y ,rad)
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
