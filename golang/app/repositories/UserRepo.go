package repositories

import (
	"app/app/models"

	"gopkg.in/mgo.v2/bson"
)

type UserRepo struct {
	Model models.User
}

func (u *UserRepo) Init() {
	u.Model = models.User{}
}

// func (u *UserRepo) ShowAll() []bson.M {
// 	models, _ := u.Model.Get()
// 	return models
// }

func (u *UserRepo) Store(body bson.M) (bson.M, error) {

	user := models.User{
		Name: body["name"].(string),
	}

	u.Model.Create(&user)

	return body, nil
}