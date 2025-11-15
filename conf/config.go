package conf

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Todolist *Config

type Config struct {
	Server *Server `yaml:"server"`
	MySQL  *MySQL  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Port      string `yaml:"port"`
	Version   string `yaml:"version"`
	JwtSecret string `yaml:"jwtSecret"`
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

type Redis struct {
	RedisDbName   string `yaml:"redisDbName"`
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisPassWord string `yaml:"redisPassWord"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	wd, _ := os.Getwd()
	viper.AddConfigPath(filepath.Join(filepath.Dir(wd), "conf"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Todolist)
	if err != nil {
		panic(err)
	}
}
