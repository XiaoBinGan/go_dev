//本段程序参考了https://www.cnblogs.com/hcy-fly/p/10826607.html博客中的内容
package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
)

const (
	consulAddress = "127.0.0.1:8500"
	localIp       = "127.0.0.1"//"your server ip"
	localPort     = 81
)

func consulRegister()  {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "337"
	registration.Name = "service337"
	registration.Port = localPort
	registration.Tags = []string{"testService"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you are visiting health check api"))
}

func main()  {
	consulRegister()
	//定义一个http接口
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe("0.0.0.0:81", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}