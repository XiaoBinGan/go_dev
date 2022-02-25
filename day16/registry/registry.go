package registry

import "context"

// Name()插件名,例如etcd
// Init(opts ...Option),初始化,里面用的选项设计模式做初始化
// Register()服务注册
// Unregister()服务反注册.例如服务端停了.注册列表销毁
// GetService 服务发现(ip port[] string)

type Registry interface{
	//plugin name
	Name() string
	Init(ctx context.Context,opts ...Option )(err error)
	Registry(ctx context.Context,service *Service)(err error)
	Unregister(ctx context.Context,service *Service)(err error)
	GetService(ctx context.Context,name string)(service *Service,err error)
}
