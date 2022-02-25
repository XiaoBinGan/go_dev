package registry

import "time"
// choose design item type



type Options struct {
	//address
	Addrs []string
	//timeout
	Timeout time.Duration
	//heartBeat
	HeartBeat int64
	// /a/b/c/XX/10.XXX
	RegistryPath string
}

//defined function type variable
type Option func(opts *Options)

func WithAddrs(addrs []string)Option  {
	return func(opts *Options) {
		opts.Addrs=addrs
	 }
}

func WithTimeout(timeout  time.Duration)Option{
	return func(opts *Options) {
		opts.Timeout=timeout
	}
}

func WithHeartBeat(heartbeat int64)Option {
	return func(opts *Options) {
		opts.HeartBeat=heartbeat
	}
}

func WithRegistryPath(registryPath string)Option  {
	return func(opts *Options) {
		opts.RegistryPath=registryPath
	}
}
