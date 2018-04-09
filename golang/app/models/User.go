
package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"app/app/config"

	"os"
)

type User struct {
	Id 					bson.ObjectId 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name				string				`json:"name" form:"name" binding:"required" bson:"name"`
}

func(u *User) Init() {
}

func(u *User) getMongoCollectionSession() *mgo.Collection {
	return config.Session.DB(os.Getenv("MONGODB_DATABASE")).C("users")
}

func(u *User) Create(model *User) {
	c := u.getMongoCollectionSession()

	model.Id = bson.NewObjectId()
	err := c.Insert(model)
	if err != nil {
		panic(err)
	}
}

// func(r *User) GetById(id string) string {
// 	newRes := Resource{}
// 	oID := bson.ObjectIdHex(id)
// 	r.getMongoCollectionSession().FindId(oID).One(&newRes)

// 	return id
// }

// func(r *Resource) Get() ([]bson.M, error) {
//     stage_match := bson.M{"$match":bson.M{"storage": "s3"}}
//     pipe := r.getMongoCollectionSession().Pipe([]bson.M{stage_match})
//     resp := []bson.M{}
// 	err := pipe.All(&resp)
    
// 	return resp, err
// }

// func(r *Resource) Create([]bson.M) {

// 	// coll := config.Session.DB(os.Getenv("MONGODB_DATABASE")).C("resources")
	
// 	//  game := Game{
//  //        Winner:       "Dave",
//  //        OfficialGame: true,
//  //        Location:     "Austin",
//  //        StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
//  //        EndTime:      time.Date(2015, time.February, 12, 05, 54, 0, 0, time.UTC),
//  //        Players: []Player{
//  //            NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
//  //            NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
//  //            NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
//  //            NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
//  //        },
//  //    }
	
// 	// if err := coll.Insert(game); err != nil {
//  //        panic(err)
//  //    }
// }