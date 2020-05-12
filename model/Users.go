package model

import "fmt"

type Users struct {
	ID            int `gorm:"primary_key"`
	UserName      string
	Passwd        string
	IsComplainted bool
}

//根据id获取user
func GetUsersById(id int) (user Users) {
	db.First(&user, id)
	return
}

//获取全部user
func GetUsers(isComplainted bool) (users []Users) {
	db.Where("is_complainted = ?", isComplainted).Find(&users)
	/*
		if isComplainted {

			db.Where("is_complainted = ?",1).Find(&users)
		} else {

			db.Where("is_complainted = ?",0).Find(&users)
		}*/
	fmt.Println(users)

	return
}

//更新
func (this *Users) Update() {
	db.Save(this)

}
