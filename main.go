package main

import (
	"ChatRoom/config"
	"ChatRoom/controller"
	"ChatRoom/model"
	"ChatRoom/router"
	"fmt"
	"time"
)

func main() {

	config.NewConfig()
	model.Setup()
	timer()
	router.StartGin()

}

func timer() {
	ticker := time.NewTicker(time.Hour * 24)
	go func() {
		for _ = range ticker.C {
			fmt.Println("ticked at %v", time.Now())
			controller.BuidRandomRoom()
		}
	}()
}
