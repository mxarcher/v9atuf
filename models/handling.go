package models

import (
	"fmt"
)

type Handling struct {
	Model
	Name      string `json:"name"`
	Algorithm string `json:"algorithm"`
	Results   string `json:"results"`
	Path      string `json:"path"`
	Comments  string `json:"comments"`
}

func GetHandlings() (handlings []Handling) {
	db.Find(&handlings)
	return
}
func GetLimitedHandlings(num int) (handlings []Handling) {
	db.Limit(num).Find(&handlings)
	return
}

func GetHandlingCount() (count int64) {
	db.Model(&Handling{}).Count(&count)
	return
}
func (h *Handling) GetHandlingByName() (handlings []Handling) {
	db.Where("name = ?", h.Name).Find(&handlings)
	return
}

func (h *Handling) AddToDB() uint {
	if h.ID > 0 {
		return 0
	}
	result := db.Create(&h)
	fmt.Println("insert id", h.ID)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0
	}
	return h.ID
}

func (h *Handling) UpdateToDB() {
	result := db.Model(&h).Updates(&h)
	if result != nil {
		panic(result)
	}
}
func (h *Handling) Delete() {
	if h.ID == 0 {
		return
	}
	err := db.Delete(&h)
	if err != nil {
		panic(err)
	}
}
