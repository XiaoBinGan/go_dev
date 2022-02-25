package common

import (
	"github.com/micro/go-micro/v2/config"
)

//Mysql config params
type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	Database  string `json:"database"`
	Port int64 `json:"port"`
}

/**get mysql config from consul
    mapping to Mysqlconfig

 */
func GetMysqlFromConsul(config config.Config,path ...string) (*MysqlConfig,error) {
	mf :=&MysqlConfig{}
	err := config.Get(path...).Scan(mf)
	if err!=nil{
		return nil, err
	}
	return mf,nil
}




