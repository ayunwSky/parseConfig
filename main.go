/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 20:40
 * @Desc:
 */

package main

import (
	// "parseConfig/viperGetItem"
	"parseConfig/viperYamlConf"
	// "parseConfig/viperItemJson"
	// "parseConfig/viperToml"
	// "parseConfig/viperTomlAndJson"
)

func main() {
	//fmt.Println("I amd main function...---------------------------")
	// viperToml.ParseToml()

	//fmt.Println("I amd main function...---------------------------")
	// viperTomlAndJson.ParseTomlAndJson()

	//fmt.Println("I amd main function...---------------------------")
	// viperItemJson.ParseJsonItem()
	
	//fmt.Println("I amd main function...---------------------------")
	// viperGetItem.ParseItems()

	viperYamlConf.ParseYamlConf()
}
