package main

import (
	"ChatRoom/config"
	"ChatRoom/model"
	"ChatRoom/router"
	"time"
)

func main() {

	config.NewConfig()
	model.Setup()
	timer()
	router.StartGin()

}

func timer() {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			//if time.Now().Minute() == 43{
			//	fmt.Println("任务 at %v", time.Now().Format("2006-01-02 15:04:05"))
			//	controller.BuidRandomRoom()
			//}
			//if time.Now().Minute() == 45 {
			//	controller.BuidRandomRoom()
			//	fmt.Println("任务 at %v", time.Now().Format("2006-01-02 15:04:05"))
			//}
			//fmt.Println("任务 at %v", time.Now().Format("2006-01-02 15:04:05"))
			//controller.BuidRandomRoom()
		}
	}()
}
