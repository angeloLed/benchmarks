package main

import (
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"app/app/config"
	"app/app"
)

func main() {

	//**************
	// Load environment
	//**************
	godotenv.Load()

	// **************
	// conntect to mongo
	// **************
	config.Connect()

	// **************
	// init Gin & Router
	// **************
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	router := app.Router{};
	router.Init(r)


	// **************
	// run
	// **************
	r.Run(":80")
}