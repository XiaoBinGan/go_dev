package main

import (
	"fmt"
	"github.com/astaxie/beego/config"//读取配置文件
)
func main() {
	conf,err :=config.NewConfig("ini","./logagent.conf")//返回一个可以对配置文件可以操作的对象
	if err!=nil {
		fmt.Println("new config failed ,err:",err)
		return
	}
	
	listenIP :=conf.Strings("server::listen_ip") //conf.Strings只有一个返回值
	fmt.Println("Listen_ip:",listenIP)

 	port,err :=conf.Int("server::port") //server::port 获取ini格式的配置文件的server下面的port参数值
	if err!=nil{
		fmt.Println("read server:port failed ,err:",err)
		return
	}
	fmt.Println("Port:",port)


	logLevel := conf.String("logs::log_level")
	fmt.Println("log_level:",logLevel)
   
	logPath :=conf.String("logs::log_path")
	fmt.Println("log_path:",logPath)


}



// // Configer defines how to get and set value from configuration raw data.
// type Configer interface {
// 	Set(key, val string) error   //support section::key type in given key when using ini type.
// 	String(key string) string    //support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
// 	Strings(key string) []string //get string slice
// 	Int(key string) (int, error)
// 	Int64(key string) (int64, error)
// 	Bool(key string) (bool, error)
// 	Float(key string) (float64, error)
// 	DefaultString(key string, defaultVal string) string      // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
// 	DefaultStrings(key string, defaultVal []string) []string //get string slice
// 	DefaultInt(key string, defaultVal int) int
// 	DefaultInt64(key string, defaultVal int64) int64
// 	DefaultBool(key string, defaultVal bool) bool
// 	DefaultFloat(key string, defaultVal float64) float64
// 	DIY(key string) (interface{}, error)
// 	GetSection(section string) (map[string]string, error)
// 	SaveConfigFile(filename string) error
// }