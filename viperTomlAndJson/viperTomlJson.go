/*
 * @ -*- Author: ayunwSky
 * @ -*- Date  : 2022/8/14 21:01
 * @ -*- Desc  :
 */

package viperTomlAndJson

import (
	"fmt"

	"github.com/spf13/viper"
)

type MySQLConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
	Title string
}

// ParseTomlAndJson 解析 toml 和 json 格式配置文件
func ParseTomlAndJson() {
	//fmt.Println("# ---------- viperTomlAndJson start ----------")
	// 定义一个 Config 结构体类型的变量 configToml

	// 读取 toml 格式配置文件使用
	var configToml Config

	vToml := viper.New()
	vToml.SetConfigName("config")
	vToml.SetConfigType("toml")
	vToml.AddConfigPath("./configs")

	if err := vToml.ReadInConfig(); err != nil {
		fmt.Printf("read toml config failed, err: %v\n", err)
		return
	}

	vToml.Unmarshal(&configToml)
	fmt.Println("read config.toml")
	fmt.Println("Config All:", configToml, "Redis:", configToml.Redis, "MySQL:", configToml.MySQL)

	fmt.Printf("Redis Host: %v\n", configToml.Redis.Host)
	fmt.Printf("MySQL Host: %v\n", configToml.MySQL.Host)
	fmt.Printf("title: %v\n", configToml.Title)

	fmt.Println("# ----------------------------------------")

	// 定义一个 Config 结构体类型的变量 configJson

	// 读取 json 格式配置文件
	var configJson Config

	vJson := viper.New()
	vJson.SetConfigName("jsconfig")
	vJson.SetConfigType("json")
	vJson.AddConfigPath("./configs")

	if err := vJson.ReadInConfig(); err != nil {
		fmt.Printf("read json config failed, err: %v\n", err)
		return
	}

	vJson.Unmarshal(&configJson)
	fmt.Println("Config All:", configJson, "Redis:", configJson.Redis, "MySQL:", configJson.MySQL)
	fmt.Printf("Redis Host: %v\n", configJson.Redis.Host)
	fmt.Printf("MySQL Host: %v\n", configJson.MySQL.Host)
	//fmt.Println("# ---------- viperTomlAndJson end ----------")
	fmt.Println()
}