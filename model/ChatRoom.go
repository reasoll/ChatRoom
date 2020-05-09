package model

type ChatRoom struct {
	ID           int `gorm:"primary_key"`
	ChatRoomName string
	CreateBy     string
	IsAnyous     bool
}

//根据id获取chatroom
func GetChatRoomById(id int) (chatroom ChatRoom) {
	db.First(&chatroom, id)
	return
}

//获取全部chatroom
func GetCharRooms() (ChatRooms []ChatRoom) {
	db.Find(&ChatRooms)
	return
}
