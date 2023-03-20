package config

import (
	"os"
	"github.com/spf13/viper"
)


func InitConfig(){
	wokdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wokdir + "/config")
	
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败失败 =>"+err.Error())
	}
}