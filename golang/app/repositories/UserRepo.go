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

func (u *UserRepo) ShowAll(filters bson.M) []bson.M {
	models, _ := u.model.Get(filters)
	data, _ := bson.Marshal(&models)
	return data
}

func (u *UserRepo) Store(body bson.M) (bson.M, error) {

	user := models.User{
		Name: body["name"].(string),
	}

	u.model.Create(&user)

	return body, nil
}