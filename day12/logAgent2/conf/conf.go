package conf

/*
	all of theme ,ready for ini config  
*/


/*AppConf config collection*/
type AppConf struct{
	KafkaConf 	`ini:"kafka"`    //ini tag one by one match
	EtcdConf	`ini:"etcd"`
}
/*KafkaConf kafka config*/
type KafkaConf struct{
	Address string `ini:"address"`
	ChanMaxSize int `ini:"chan_max_size"`
}

/*EtcdConf config etcd client*/
type EtcdConf struct{
	Address string `ini:"address"`
	Key string `ini:"collect_log_key"`
	TimeOut int `ini:"timeout"`
}

/*			unused  			*/

/*TaillogConf tail file config*/
type TaillogConf struct{
	FileName string `ini:"filename"`
}