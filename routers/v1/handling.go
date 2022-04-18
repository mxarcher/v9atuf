package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/models"
	"net/http"
)

func GetHandling(ctx *gin.Context) {

	data := make(map[string]interface{})
	data["Handling"] = models.GetLimitedHandlings(20)
	data["count"] = models.GetHandlingCount()
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func AddHandling(ctx *gin.Context) {
	json := info[models.Handling]{}
	err := ctx.ShouldBind(&json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Model.AddToDB()

	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Operation: "add",
	}
	log.AddToDB()

	ctx.JSON(http.StatusOK, gin.H{
		"data": json,
	})

}

func UpdateHandling(ctx *gin.Context) {
	json := info[models.Handling]{}
	err := ctx.ShouldBind(&json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Model.UpdateToDB()

	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Operation: "update",
	}
	log.AddToDB()

	fmt.Printf("%v\n", json)
	ctx.JSON(http.StatusOK, gin.H{
		"data": json,
	})
}

func DeleteHandling(ctx *gin.Context) {
	json := info[models.Handling]{}
	err := ctx.ShouldBind(json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.Model.Delete()
	log := models.Log{
		Tid:       json.Model.ID,
		Name:      json.Operator,
		Operation: "delete",
	}
	log.AddToDB()

}
