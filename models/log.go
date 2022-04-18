package models

import (
	"fmt"
)

type Log struct {
	Model
	Name      string `json:"name"`
	Operation string `json:"operation"`
	Table     string `json:"table"`
	Tid       uint   `json:"tid"`
}

func GetLimitedLogs(num int) (logs []Log) {
	db.Limit(num).Find(&logs)
	return
}
func (l *Log) UpdateToDB() {
	result := db.Model(&l).Updates(&l)
	if result != nil {
		panic(result)
	}
}
func (l *Log) AddToDB() uint {
	if l.ID > 0 {
		return 0
	}
	result := db.Create(&l)
	fmt.Println("insert id", l.ID)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0
	}
	return l.ID
}
func (l *Log) Delete() {
	if l.ID == 0 {
		return
	}
	db.Delete(&l)
}
