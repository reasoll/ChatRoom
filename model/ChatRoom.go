package model

type ChatRoom struct {
	ID             int `gorm:"primary_key"`
	ChatRoomName   string
	CreateBy       string
	IsAnyous       bool
	WhoComplainted string
}

//根据id获取chatroom
func GetChatRoomById(id int) (chatroom ChatRoom) {
	db.First(&chatroom, id)
	return
}

//获取全部chatroom
func GetCharRooms() (ChatRooms []ChatRoom) {
	db.Order("id desc").Find(&ChatRooms)
	return
}

func (this *ChatRoom) CreateRoom() error {
	return db.Create(this).Error
}
