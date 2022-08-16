/*
 * @ -*- Author: ayunwSky
 * @ -*- Date: 2022/8/14 21:32
 * @ -*- Desc:
 */

package viperGetItem

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/viper"
)

// typeAssert 断言
// func typeAssert(a interface{}) interface{}{
// 	fmt.Printf("%T\n",a)

// 	switch v := a.(type){
// 	case string:
// 		return v
// 	case int:
// 		return v
// 	default:
// 		fmt.Printf("unsupport type!")
// 	}
// 	return
// }

// ParseItems 读取配置项的值
func ParseItems() {
	//fmt.Println("# ---------- viperGetItem start ----------")

	// 设置要解析的配置文件名
	viper.SetConfigName("item")
	// 设置要解析的配置文件格式
	viper.SetConfigType("json")
	// 设置配置文件所在路径，如果是当前路径，就写成 "."
	viper.AddConfigPath("./configs")

	// 根据上面的配置加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("load item config file failed, err:%v\n", err)
		return
	}

	mysqlHost := viper.Get("mysql.host")
	mysqlUsername := viper.Get("mysql.username")
	mysqlPort := viper.Get("mysql.port")
	mysqlPortsSlice := viper.GetIntSlice("mysql.ports")

	mysqlMetricsPort := viper.GetInt("mysql.metrics.port")
	redisConf := viper.Get("redis")

	if viper.IsSet("mysql.host") {
		fmt.Println("[IsSet()] mysql.host is set")
	} else {
		fmt.Println("[IsSet()] mysql.host is not set")
	}

	fmt.Println("mysqlHost:", mysqlHost, ", mysqlUsername:", mysqlUsername, ", mysqlPort:", mysqlPort)
	fmt.Println("mysqlPorts:", mysqlPortsSlice)
	fmt.Println("mysqlMetricsPort:", mysqlMetricsPort)
	fmt.Println("redisConf:", redisConf)

	fmt.Println()
	fmt.Println("---------- mysqlStringMap ----------")
	mysqlStringMap := viper.GetStringMap("mysql")

	fmt.Println("mysqlStringMap:", mysqlStringMap)
	fmt.Println("mysqlStringMapUsername:", mysqlStringMap["username"])

	// v := typeAssert(mysqlStringMap["metrics.host"])
	// fmt.Println(v)

	
	// jsonData :=[]byte(mysqlStringMap)

	// var v interface{}
	// json.Unmarshal(jsonData, &v)
	// data := v.(map[string]interface{})

	// for k, v := range data{
	// 	switch v := v.(type) {
	// 	case string:
	// 		fmt.Println(k,v , "(string)")
	// 	case float64:
	// 		fmt.Println(k,v , "(float64)")
	// 	case []interface{}:
	// 		fmt.Println(k,"(array)")
	// 		for i, u := range v{
	// 			fmt.Println("  ",i,u)
	// 		}
	// 	default:
	// 		fmt.Println(k,v,"(unknow)")
	// 	}
	// }
	
	fmt.Println("mysqlStringMapHost:", mysqlStringMap["metrics.host"])
	fmt.Println("mysqlStringMapMetricsPort:", mysqlStringMap["metrics.port"])
}
