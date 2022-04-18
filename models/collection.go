package models

import (
	"fmt"
)

type Collection struct {
	Model
	// 用下面的两个声明从属外键关系
	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
	//
	Name     string `json:"name"`
	Program  string `json:"program"`
	DataSet  string `json:"dataset"`
	Path     string `json:"path"`
	Comments string `json:"comments"`
}

func GetCollections() (collections []Collection) {
	db.Find(&collections)
	return
}
func GetLimitedCollections(num int) (collections []Collection) {
	db.Limit(num).Find(&collections)
	return
}

func GetCollectionCount() (count int64) {
	db.Model(&Collection{}).Count(&count)
	return
}
func (c *Collection) GetCollectionByName() (collections []Collection) {
	db.Where("name = ?", c.Name).Find(&collections)
	return
}
func (c *Collection) AddToDB() uint {
	if c.ID > 0 {
		return 0
	}
	result := db.Create(&c)
	fmt.Println("insert id", c.ID)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0
	}
	return c.ID
}

func (c *Collection) UpdateToDB() {
	result := db.Model(&c).Updates(&c)
	if result != nil {
		panic(result)
	}
}
func (c *Collection) Delete() {
	if c.ID == 0 {
		return
	}
	db.Delete(&c)
}
