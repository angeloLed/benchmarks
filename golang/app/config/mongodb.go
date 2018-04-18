package config

import (
	"fmt"
	"os"
	"gopkg.in/mgo.v2"
)

var (
	
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

var MongoDbConnected = make(chan bool)

// Connect connects to mongodb
func Connect() {
	uri := ""

	if os.Getenv("MONGODB_CONNECTION_STRING") != "" {
		uri = os.Getenv("MONGODB_CONNECTION_STRING") 
	} else {
		uri = "mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT") + "/" + os.Getenv("MONGODB_DATABASE")
	}

	fmt.Printf("Connecting to mongoDb server...\n")
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)


	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)

	Session = s
	Mongo = mongo
}