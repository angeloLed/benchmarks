package services

import (
	"app/app/repositories"

	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	repo repositories.UserRepo
}

func (u *UserService) Init() {
	u.repo = repositories.UserRepo{}
}

// func (r *UserService) ShowMany() []bson.M {
// 	models, _ := r.Model.Get()
// 	return models
// }

func (u *UserService) Store(body bson.M) (bson.M, error) {
	return u.repo.Store(body)
}