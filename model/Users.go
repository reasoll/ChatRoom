package model

type Users struct {
	ID				int
	UserName		string
	Passwd			string
	IsComplained	bool
}


//根据id获取chatroom
func GetUsersById(id int)(user Users) {
	db.First(&user,id)
	return
}

//获取全部chatroom
func GetUsers() (users []Users) {
	db.Find(&users)
	return
}

