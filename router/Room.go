package router

import "ChatRoom/controller"

func RoomRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/roomlist11",controller.GetRoomList)
	r.GET("/test")


}
