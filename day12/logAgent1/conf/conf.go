package conf

/*
	all of theme ,ready for ini config  
*/


/*AppConf config collection*/
type AppConf struct{
	KafkaConf 	`ini:"kafka"`    //ini tag one by one match
	TaillogConf	`ini:"taillog"`
}
/*KafkaConf kafka config*/
type KafkaConf struct{
	Arrdress string `ini:"address"`
	Topic    string `ini:"topic"`
}
/*TaillogConf tail file config*/
type TaillogConf struct{
	FileName string `ini:"filename"`
}