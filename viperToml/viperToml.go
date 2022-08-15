/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 20:39
 * @Desc:
 */

package viperToml

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

func ParseToml() {
	fmt.Println("# ---------- viperToml start ----------")
	// 定义一个 Config 结构体类型的变量 config
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper read config failed, err: %v\n", err)
		return
	}

	viper.Unmarshal(&config)

	fmt.Println("Config: ", config, "Redis: ", config.Redis, "MySQL: ", config.MySQL)
	fmt.Printf("Redis Host: %v\n", config.Redis.Host)
	fmt.Printf("title: %v\n", config.Title)

	fmt.Println("# ---------- viperToml end ----------")
	fmt.Println()
}
