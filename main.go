package main

import (
	"ChatRoom/config"
	"ChatRoom/model"
	"ChatRoom/router"
)

func main() {

	config.NewConfig()
	model.Setup()
	router.StartGin()

}
