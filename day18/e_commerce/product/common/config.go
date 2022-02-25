package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string,port int64,prefix string)(config.Config,error){
  //add config center
  //config center use consul's key /value mode
	source := consul.NewSource(
		//config center address
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//set prefix   the default /micro/config
		consul.WithPrefix(prefix),
		// is remove prefix there set true, it means not set prefix still can be got the config
		consul.StripPrefix(true),
	)
	//init config
	conf, err := config.NewConfig()
	if err!=nil{
		return conf, err
	}
	err = conf.Load(source)
	return conf,err
}