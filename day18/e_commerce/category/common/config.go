package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)
/**Get Consul Config
   @params{host string}
   @params{port int64}
   @params{prefix string}
   return config.Config, error
 */
func GetConsulConfig(host string,port int64,prefix string)(config.Config,error)  {
	consulSource := consul.NewSource(
		//config conf_center address
		//because with address the receive parameter is of type string so the int64 type is converted to type String
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//set prefix,don't set default prefix
		consul.WithPrefix(prefix),
		//is remove prefix,set to true here,Indicates that the corresponding configuration can be obtained directly without a prefix
		consul.StripPrefix(true),
	)
	//init conf
	config, err := config.NewConfig()
	if err!=nil{
		return config, err
	}
	//lode config
	err = config.Load(consulSource)
	return config,err
}
