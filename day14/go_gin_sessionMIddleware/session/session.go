package session
//定义接口只负责定义规范 具体实现有具体的文件
type Session interface {
	Set(key string,value interface{})error
	Get(key string)(interface{},error)
	Del(key string)error
	Save()error
}
