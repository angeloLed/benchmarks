package services

import (
	"app/app/repositories"
	"gopkg.in/mgo.v2/bson"
	"app/app/models"
)

type UserService struct {
	repo repositories.UserRepo
}

func (u *UserService) Init() {
	u.repo = repositories.UserRepo{}
}

func (u *UserService) ShowMany(filters bson.M) ([]models.User) {
	models := u.repo.ShowAll(filters)
	return models
}

func (u *UserService) GetAllUserHasHeatZone(x int, y int, radius int) ([]string) {

	filters := bson.M{}

	minx := x - radius
	maxx := y + radius
	filters["x"] = bson.M{
		"$gte": minx,
		"$lte": maxx,
	}

	miny := x - radius
	maxy := y + radius
	filters["y"] = bson.M{ 
		"$gte": miny,
		"$lte": maxy,
	}

	models := u.ShowMany(filters)
	var users []string

	for _, element := range models {
		if ! u.contains(users,element.User) {
			users = append(users, element.User)
		}
	}

	return users
}

func (u *UserService) contains(sliceValues []string, search string) bool {
    for _, v := range sliceValues {
        if v == search {
            return true
        }
    }
    return false
}

func (u *UserService) Store(body bson.M) (bson.M, error) {
	return u.repo.Store(body)
}