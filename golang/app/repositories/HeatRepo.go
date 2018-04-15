package repositories

import (
	"app/app/models"
	"gopkg.in/mgo.v2/bson"
)

type HeatRepo struct {
	model models.Heat
}

func (u *HeatRepo) Init() {
	u.model = models.Heat{}
}

func (u *HeatRepo) GetAllUserHasHeatZone(filters bson.M) []models.Heat {
	models, _ := u.model.Get(filters)
	return models
}

func (u *HeatRepo) ShowAll(filters bson.M) []models.Heat {
	models, _ := u.model.Get(filters)
	return models
}

func (u *HeatRepo) Store(body bson.M) (bson.M, error) {

	heat := models.Heat{
		User: body["user"].(string),
		X: int(body["x"].(float64)),
		Y: int(body["y"].(float64)),
	}

	u.model.Create(&heat)

	return body, nil
}