package controller

import (
	"ChatRoom/model"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-broadcast"
	"github.com/gin-gonic/gin"
	"html"
	"io"
	"net/http"
	"strings"
	"time"
)

var RoomChannels = make(map[string]broadcast.Broadcaster)

func OpenListener(roomid string) chan interface{} {
	listener := make(chan interface{})
	Room(roomid).Register(listener)
	return listener
}

func CloseListener(roomid string, listener chan interface{}) {
	Room(roomid).Unregister(listener)
	close(listener)
}

func Room(roomid string) broadcast.Broadcaster {
	b, ok := RoomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		RoomChannels[roomid] = b
	}
	return b
}



func RateLimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(http.StatusServiceUnavailable, "you were automatically banned :)")
	}
}

func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/room/hn")
}

func RoomList(c *gin.Context) {
	c.HTML(http.StatusOK, "room_login.templ.html", gin.H{
		"roomid":    "testRoom",
		"nick":      "testNice",
		"timestamp": time.Now().Unix(),
	})
}

func RoomGET(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	if len(nick) < 2 {
		nick = ""
	}
	if len(nick) > 13 {
		nick = nick[0:12] + "..."
	}
	c.HTML(http.StatusOK, "room_login.templ.html", gin.H{
		"roomid":    roomid,
		"nick":      nick,
		"timestamp": time.Now().Unix(),
	})

}

func RoomPOST(c *gin.Context) {
	roomid := c.Param("roomid")
	nick := c.Query("nick")
	message := c.PostForm("message")
	message = strings.TrimSpace(message)

	validMessage := len(message) > 1 && len(message) < 200
	validNick := len(nick) > 1 && len(nick) < 14
	if !validMessage || !validNick {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "the message or nickname is too long",
		})
		return
	}

	post := gin.H{
		"nick":    html.EscapeString(nick),
		"message": html.EscapeString(message),
	}
	messages.Add("inbound", 1)
	Room(roomid).Submit(post)
	c.JSON(http.StatusOK, post)
}

func StreamRoom(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := OpenListener(roomid)
	ticker := time.NewTicker(1 * time.Second)
	users.Add("connected", 1)
	defer func() {
		CloseListener(roomid, listener)
		ticker.Stop()
		users.Add("disconnected", 1)
	}()

	c.Stream(func(w io.Writer) bool {
		select {
		case msg := <-listener:
			messages.Add("outbound", 1)
			c.SSEvent("message", msg)
		case <-ticker.C:
			c.SSEvent("stats", Stats())
		}
		return true
	})
}



//获取所有聊天室
func GetRoomList(c *gin.Context) {
	roomList :=  model.GetCharRooms()
	jsons, _ := json.Marshal(roomList)
	fmt.Print(string(jsons))
	c.HTML(http.StatusOK, "room_list.templ.html", string(jsons))



}
