package main

import "fmt"

//
//import (
//	"fmt"
//	viper "github.com/spf13/viper"
//	 _ "github.com/spf13/viper/remote"
//)
//func main() {
////	viper.SetDefault("ContentDir", "content")
////	viper.SetDefault("LayoutDir", "layouts")
////	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
//	viper.SetConfigFile("./index") // 指定配置文件路径
////	viper.SetConfigName("config") // 配置文件名称(无扩展名)
//	viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
////	viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
////	viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
////	viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
//	err := viper.ReadInConfig() // 查找并读取配置文件
//	if err != nil { // 处理读取配置文件的错误
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//	get := viper.Get("name") // 这里会得到 "steve"
//	fmt.Println(get)
//	viper.Set("name", "name")
//	get = viper.Get("name") // 这里会得到 "steve"
//	fmt.Println(get)
//
////	if err := viper.ReadInConfig(); err != nil {
////		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
////			// 配置文件未找到错误；如果需要可以忽略
////			fmt.Printf("%v\n","fail")
////		} else {
////			// 配置文件被找到，但产生了另外的错误
////			fmt.Printf("%v\n","fail2")
////		}
////	}
////			// 配置文件找到并成功解析
////			fmt.Printf("%v\n","success")
//
//	//viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")
//
//	//// 任何需要将此配置添加到程序中的方法。
//	//var yamlExample = []byte(`
//	//	Hacker: true
//	//	name: steve
//	//	hobbies:
//	//	- skateboarding
//	//	- snowboarding
//	//	- go
//	//	clothing:
//	//	  jacket: leather
//	//	  trousers: denim
//	//	age: 35
//	//	eyes : brown
//	//	beard: true
//	//	`)
//	//
//	//viper.ReadConfig(bytes.NewBuffer(yamlExample))
//	//
//	//get := viper.Get("name") // 这里会得到 "steve"
//	//fmt.Println(get)
//}
func main()  {
	var n string
	fmt.Println(n)
	var a [3]int
	var b =  [...]int{1,2,3,}
	var c =  [...]int{2:3,1:2}
	fmt.Println(a,"\n",b,"\n",c,"\n")
}