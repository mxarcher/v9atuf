package routers

import (
	"github.com/gin-gonic/gin"
	"learn/pkg/setting"
	v1 "learn/routers/v1"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	{
		api.GET("/user", v1.GetUsers)
		api.POST("/user", v1.AddUser)
		api.PUT("/user", v1.UpdateUser)
		api.DELETE("/user", v1.DeleteUser)

		api.GET("/collection", v1.GetCollection)
		api.POST("/collection", v1.AddCollection)
		api.PUT("/collection", v1.UpdateCollection)
		api.DELETE("/collection", v1.DeleteCollection)

		api.GET("/handling", v1.GetHandling)
		api.POST("/handling", v1.AddHandling)
		api.PUT("/handling", v1.UpdateHandling)
		api.DELETE("/handling", v1.DeleteHandling)

		api.GET("/log", v1.GetLogs)
	}
	err := r.SetTrustedProxies(setting.WhiteList)
	if err != nil {
		return nil
	}
	return r
}
