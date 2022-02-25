package registry

import (
	"context"
	"fmt"
	"sync"
)

//plugin  manager class
//use a manager map ,string of the key ,value is the interface object of the  Registry
//defined to use ,defined plugin
//implement center of the registry,provide system use .

//defined manager struct

type PluginMgr struct {
	//map maintain all plugin
	plugin map[string]Registry
	lock sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugin: make(map[string]Registry),
	}
)

//public plugin registry
func RegistryPlugin(registry Registry)(err error)  {
	return pluginMgr.register(registry)
}

//private  registry plugin

func(p *PluginMgr)register(plugin Registry)(err error){
	p.lock.Lock()
	defer p.lock.Unlock()
	_,ok := p.plugin[plugin.Name()]
	if !ok{
		fmt.Errorf("register plugin exit")
		return
	}
	p.plugin[plugin.Name()]=plugin
	return
}

// public init Registry
func InitRegistry(ctx context.Context,name string,opt Option)(registry Registry,err error){
	return pluginMgr.initRegistry(ctx,name,opt)
}

//private initRegistry
func(p *PluginMgr)initRegistry(ctx context.Context,name string,opt Option)(registry Registry,err error){
	p.lock.Lock()
	defer p.lock.Unlock()
	plugin,ok := p.plugin[name]
	if !ok{
		fmt.Errorf("plugin %s is not exist",name)
		return
	}
	registry =plugin
	err = plugin.Init(ctx, opt)
	return
}
