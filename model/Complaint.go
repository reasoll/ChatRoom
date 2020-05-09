package model

type Complaint struct {
	ID				int		`gorm:"primary_key"`
	UserName		string
	WhoComplainted	string
	Content 		string
	RoomID			int
}

//根据id获取chatroom
func GetComplainById(id int)(complaint Complaint) {
	db.First(&complaint,id)
	return
}

//获取全部chatroom
func GetComplaints(roomID int) (complaints []Complaint) {
	db.Where("room_id = ?",roomID).Find(&complaints)
	return
}

