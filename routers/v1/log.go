package v1

import (
	"github.com/gin-gonic/gin"
	"learn/models"
	"net/http"
)

func GetLogs(ctx *gin.Context) {

	data := make(map[string]interface{})
	data["logs"] = models.GetLimitedLogs(50)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
