package router

import "ChatRoom/controller"

func RoomRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/createroom", controller.CreateRoomPre)
	r.POST("/createroom", controller.CreateRoom)

}
