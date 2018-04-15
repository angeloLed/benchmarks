
package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"app/app/config"

	"os"
)

type Heat struct {
	Id		bson.ObjectId	`json:"_id,omitempty" bson:"_id,omitempty"`
	User 	string 			`bson:"user" json:"user"`
	X 		int 			`bson:"x" json:"x"`
	Y 		int 			`bson:"y" json:"y"`
}

func(u *Heat) Init() {
}

func(u *Heat) getMongoCollectionSession() *mgo.Collection {
	return config.Session.DB(os.Getenv("MONGODB_DATABASE")).C("usersheats")
}

func(u *Heat) Create(model *Heat) {
	c := u.getMongoCollectionSession()

	model.Id = bson.NewObjectId()
	err := c.Insert(model)
	if err != nil {
		panic(err)
	}
}

func(u *Heat) Get(query bson.M) ([]Heat, error) {
	results := []Heat{}
	c := u.getMongoCollectionSession()
	err := c.Find(query).All(&results)

	return results, err
}