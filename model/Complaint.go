package model

type Complaint struct {
	ID             int `gorm:"primary_key"`
	Username       string
	WhoComplainted string
	Content        string
	RoomID         int
}

//根据id获取Complaints
func GetComplainById(id int) (complaint Complaint) {
	db.First(&complaint, id)
	return
}

//获取全部Complaints
func GetComplaints(roomID int) (complaints []Complaint) {
	db.Where("room_id = ?", roomID).Find(&complaints)
	return
}

//create
func (this *Complaint) Create() {
	db.Create(this)
}
