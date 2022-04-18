package setting

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ServerPort int
	RunMode    string
	DBHost     string
	DBUser     string
	DBPasswd   string
	DBName     string
	DBPort     int
	WhiteList  []string
)

func init() {
	viper.SetConfigFile("conf/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fail to parse 'conf/config.yml")
		return
	}
	RunMode = viper.GetString("mode")

	ServerPort = viper.GetInt("server.port")

	// 读取数据库信息
	DBHost = viper.GetString("db.host")
	DBName = viper.GetString("db.name")
	DBUser = viper.GetString("db.user")
	DBPasswd = viper.GetString("db.password")
	DBPort = viper.GetInt("db.port")
	WhiteList = viper.GetStringSlice("whitelist")
}
