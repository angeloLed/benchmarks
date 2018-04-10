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

func (u *UserService) ShowMany(filters bson.M) []bson.M {
	models := u.repo.ShowAll(filters)
	return models
}

func (u *UserService) GetAllUserHasHeatZone(parameters bson.M) []bson.M {

	var filters bson.M{}

	if parameters["x"] != nil {
		min := parameters["x"].(int) - parameters["radius"].(int)
		max := parameters["x"].(int) + parameters["radius"].(int)
		filters["x"] = bson.M{ "$gte": min,  "$lte": max }
	}

	if parameters["y"] != nil {
		min := parameters["y"].(int) - parameters["radius"].(int)
		max := parameters["y"].(int) + parameters["radius"].(int)
		filters["y"] = bson.M{ "$gte": min,  "$lte": max }
	}

	if parameters["user"] != nil {
		filters["user"] = parameters["radius"]
	}

	models := u.ShowMany(filters)
	return models
}

func (u *UserService) Store(body bson.M) (bson.M, error) {
	return u.repo.Store(body)
}