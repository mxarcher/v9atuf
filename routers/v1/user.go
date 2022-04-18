package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/models"
	"net/http"
)

func GetUsers(ctx *gin.Context) {

	data := make(map[string]interface{})
	data["users"] = models.GetLimitedUsers(20)
	data["count"] = models.GetUserCount()
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func AddUser(ctx *gin.Context) {
	json := info[models.User]{}
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(json)
	json.Model.AddToDB()

	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Table:     "user",
		Operation: "add",
	}
	log.AddToDB()

	ctx.JSON(http.StatusOK, gin.H{
		"data": json,
	})

}

func UpdateUser(ctx *gin.Context) {
	json := info[models.User]{}
	err := ctx.ShouldBind(&json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Model.UpdateToDB()

	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Table:     "user",
		Operation: "update",
	}
	log.AddToDB()

	fmt.Printf("%v\n", json)
	ctx.JSON(http.StatusOK, gin.H{
		"data": json,
	})
}

func DeleteUser(ctx *gin.Context) {
	json := info[models.User]{}
	err := ctx.ShouldBind(json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Model.Delete()
	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Table:     "user",
		Operation: "delete",
	}
	log.AddToDB()

}
