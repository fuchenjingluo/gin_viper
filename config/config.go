package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var AppConfig *Config

// 读取配置文件yml
func InitConfig() {
	viper.SetConfigName("config")       //配置文件名称（无扩展名）
	viper.SetConfigType("yaml")         //扩展名
	viper.AddConfigPath("./config.yml") //文件路径
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil { //将配置文件中的内容反序列化到AppConfig中
		panic(err)
	}
	//初始化数据库
	InitDB()
}
