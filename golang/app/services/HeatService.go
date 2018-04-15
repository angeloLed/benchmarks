package services

import (
	"app/app/repositories"
	"gopkg.in/mgo.v2/bson"
	"app/app/models"
)

type HeatService struct {
	repo repositories.HeatRepo
}

func (u *HeatService) Init() {
	u.repo = repositories.HeatRepo{}
}

func (u *HeatService) ShowMany(filters bson.M) ([]models.Heat) {
	models := u.repo.ShowAll(filters)
	return models
}

func (u *HeatService) GetAllUserHasHeatZone(x int, y int, radius int) ([]string) {

	filters := bson.M{}

	minx := x - radius
	maxx := x + radius
	filters["x"] = bson.M{
		"$gte": minx,
		"$lte": maxx,
	}

	miny := y - radius
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

func (u *HeatService) contains(sliceValues []string, search string) bool {
    for _, v := range sliceValues {
        if v == search {
            return true
        }
    }
    return false
}

func (u *HeatService) Store(body bson.M) (bson.M, error) {
	return u.repo.Store(body)
}