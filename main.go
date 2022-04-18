package main

import (
	"fmt"
	"learn/pkg/setting"
	"learn/routers"
)

func main() {
	router := routers.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", setting.ServerPort))
	if err != nil {
		return
	}
}
