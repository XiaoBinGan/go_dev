package session

import (
	"fmt"
	"sync"
)

//对象
//MemorySession对象（字段：sessionid ,存取kv的map ，读写锁）
//定义构造函数，获取新的对象
//实现Session interface
//Set()
//Get()
//Del()
//Save()
type MemorySession struct {
	sessionId string
	data map[string]interface{}
	rwLock sync.RWMutex
}

func NewMemorySession(id string)*MemorySession  {
	 return &MemorySession{
		sessionId: id,
		data: make(map[string]interface{},20),
	}
}
func (m *MemorySession)Set(key string,value interface{})error{
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.data[key]=value
	return nil
}
func (m *MemorySession)Get(key string)(value interface{},err error){
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	value,ok :=m.data[key]
	if !ok{
		fmt.Printf("key:%v not in session\n",key)
		return
	}
	return
}
func (m *MemorySession)Del(key string)(err error){
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data,key)
	return
}
func (m *MemorySession)Save()(err error){
	return
}