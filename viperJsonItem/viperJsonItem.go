/*
 * @ -*- Author: ayunwSky
 * @ -*- Date  : 2022/8/15 19:53
 * @ -*- Desc  :
 */

package viperItemJson

import (
	"fmt"
	"github.com/spf13/viper"
)

// 用 viper 获取 json 配置文件

type MySQLConfig struct {
	Port     int      `json:"port"`
	Host     string   `json:"host"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Ports    []string `json:"ports"`
	Metrics  Metrics  `json:"metrics"`
}

type Metrics struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	MySQL MySQLConfig
	Metrics Metrics
	Redis string
}

// ParseJsonItem 从 item.json 文件解析数据
func ParseJsonItem() {
	//fmt.Println("# ---------- parse item.json start ----------")

	var itemConfigJson Config

	vJson := viper.New()
	vJson.SetConfigName("item")
	vJson.SetConfigType("json")
	vJson.AddConfigPath("./configs")

	if err := vJson.ReadInConfig(); err != nil {
		fmt.Printf("read in item.json file failed, err: %v\n", err)
		return
	}

	vJson.Unmarshal(&itemConfigJson)
	fmt.Println("start parse item.json")
	fmt.Println("config all:", itemConfigJson)

	fmt.Println("# ---------- MySQL Info ----------")
	fmt.Println("MySQL:", itemConfigJson.MySQL)
	fmt.Println("MySQL.Port:", itemConfigJson.MySQL.Port)
	fmt.Println("MySQL.Host:", itemConfigJson.MySQL.Host)
	fmt.Println("MySQL.Username:", itemConfigJson.MySQL.Username)
	fmt.Println("MySQL.Password:", itemConfigJson.MySQL.Password)
	fmt.Println("MySQL.Ports:", itemConfigJson.MySQL.Ports)

	fmt.Println()
	fmt.Println("# ---------- Metrics Info ----------")
	fmt.Println("Metrics:", itemConfigJson.Metrics)
	fmt.Println("Metrics.Host:", itemConfigJson.Metrics.Host)
	fmt.Println("Metrics.Port:", itemConfigJson.Metrics.Port)

	fmt.Println()
	fmt.Println("# ---------- Redis Info ----------")
	fmt.Println("Redis:", itemConfigJson.Redis)
}