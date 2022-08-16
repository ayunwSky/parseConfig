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
	// 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	vJson.SetConfigName("item")
	// 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	vJson.SetConfigType("json")
	// 查找配置文件所在路径
	vJson.AddConfigPath("./configs")
	// 多次调用AddConfigPath，可以添加多个搜索路径
	// viper.AddConfigPath("/etc/appname/")
	// 还可以在工作目录中搜索配置文件
	// viper.AddConfigPath(".")

	// 根据上面的配置加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		// fmt.Printf("load item config file failed, err:%v\n", err)
		panic(fmt.Errorf("load item config file failed, err:%v\n", err))
		return
	}

	vJson.Unmarshal(&itemConfigJson)
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