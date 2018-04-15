package controllers

import (
	"net/http"

	"app/app/services"
	"strconv"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

type HeatController struct {
	Service services.HeatService
}

func GetInstance() HeatController {
	service := services.HeatService{}
	service.Init()
	return HeatController{Service: service}
}

func (u *HeatController) GetAll(c *gin.Context) {

	var filters bson.M //implementation of the filters in another spacetime

	heats := u.Service.ShowMany(filters)
	c.JSON(http.StatusOK, gin.H{
		"data": heats,
	})
}

func (u *HeatController) GetAllUserHasHeatZone(c *gin.Context) {
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


func(u *HeatController) Store(c *gin.Context) {
	
	body := bson.M{}
	c.BindJSON(&body)

	heat, err := u.Service.Store(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":  heat,
	})
}
