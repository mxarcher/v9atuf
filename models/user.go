package models

import (
	"fmt"
)

type User struct {
	Model
	Name     string  `json:"name"`
	Sex      bool    `json:"sex"`
	Height   float32 `json:"height" gorm:"precision:5;scale:2"`
	Weight   float32 `json:"weight" gorm:"precision:5;scale:2"`
	Comments string  `json:"comments"`
}

func GetUsers() (users []User) {
	db.Find(&users)
	return
}
func GetLimitedUsers(num int) (users []User) {
	db.Limit(num).Find(&users)
	return
}

func GetUserCount() (count int64) {
	db.Model(&User{}).Count(&count)
	return
}
func GetUserByName(name string) (users []User) {
	db.Where("name = ?", name).Find(&users)
	return
}

func (user *User) AddToDB() uint {
	if user.ID > 0 {
		return 0
	}
	result := db.Create(&user)
	fmt.Println("insert id", user.ID)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0
	}
	return user.ID
}

func (user *User) UpdateToDB() {
	db.Updates(&user)
}

func (user *User) Delete() {
	if user.ID == 0 {
		return
	}
	db.Delete(&user)
}
