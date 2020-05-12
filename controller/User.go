package controller

import (
	"ChatRoom/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

//
func GetUserList(c *gin.Context) {
	dd := c.Param("dd")
	fmt.Println("dd:", dd)
	flg, _ := strconv.ParseBool(dd)
	fmt.Println("flg:", flg)
	users := model.GetUsers(flg)
	c.JSON(200, users)

}

//每天新建一个聊天室
func BuidRandomRoom() {

	user, flg := GetRandomUser()
	if flg {
		ChatRoomName := time.Now().Format("2006-01-02") + " 吐糟聊天室"
		fmt.Print(ChatRoomName)
		WhoComplainted := user.UserName
		CreateBy := ""
		IsAnyous := false

		chatRoom := model.ChatRoom{0, ChatRoomName, CreateBy, IsAnyous, WhoComplainted}
		chatRoom.CreateRoom()
		user.IsComplainted = true
		user.Update()
	}

}

//随机获取一个用户
func GetRandomUser() (user model.Users, err bool) {
	users := model.GetUsers(false)
	if len(users) < 1 {
		return model.Users{}, false

	}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(users))
	return users[i], true
}
