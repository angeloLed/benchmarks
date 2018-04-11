package repositories

import (
	"app/app/models"
	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	model models.User
}

func (u *UserRepo) Init() {
	u.model = models.User{}
}

func (u *UserRepo) GetAllUserHasHeatZone(filters bson.M) []models.User {
	models, _ := u.model.Get(filters)
	return models
}

func (u *UserRepo) ShowAll(filters bson.M) []models.User {
	models, _ := u.model.Get(filters)
	return models
}

func (u *UserRepo) Store(body bson.M) (bson.M, error) {

	user := models.User{
		User: body["user"].(string),
		X: int(body["x"].(float64)),
		Y: int(body["y"].(float64)),
	}

	u.model.Create(&user)

	return body, nil
}