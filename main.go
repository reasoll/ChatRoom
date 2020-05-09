package main

import (
	"ChatRoom/config"
	"ChatRoom/controller"
	"ChatRoom/model"
	"ChatRoom/router"
	"fmt"
	"runtime"
)

func main() {

	config.NewConfig()
	model.Setup()
	ConfigRuntime()
	StartWorkers()
	router.StartGin()


	
}


// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go controller.StatsWorker()
}



