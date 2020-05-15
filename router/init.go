package router

import (
	"ChatRoom/controller"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var Router *gin.Engine

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(controller.RateLimit, gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", controller.Index)
	router.GET("/roomlist", controller.GetRoomList)
	router.GET("/room/:roomid", controller.RoomGET)
	router.POST("/room-post/:roomid/:whoComplainted", controller.RoomPOST)
	router.GET("/stream/:roomid", controller.StreamRoom)
	router.GET("/users/:dd", controller.GetUserList)
	router.GET("/testajax", controller.Testajax)
	router.GET("/addjson", controller.Addjson)
	router.GET("/download", controller.DownloadAll)

	Router = router
	RoomRouter("roomm")

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
