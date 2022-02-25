package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Database string `json:"database"`
	Port int64  `json:"port"`
}
/**
   name GetMysqlFromConsul
   init MysqlConfig struct
   user config.Get(path...).Scan(mysqlconfig)
   return mysqlConfig
 */
func GetMysqlFromConsul(config config.Config,path ...string)*MysqlConfig {
	mysqlConfig :=&MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig

}

