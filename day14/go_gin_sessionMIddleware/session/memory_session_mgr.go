package session

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
)

//对象
//MemorySessionMgr对象（字段：sessionidMap ,读写锁）
//定义构造函数，获取新的对象
//实现SessionMgr interface
//Init()
//CreateSession()
//Get()
/**
------------------MemorySessionMgr-----------------------------
|sessionId|map[string]interface{}				| 备注 		  |
---------------------------------------------------------------
|_UUId0xk1|map[session1]{"name":123},....repeat |MemorySession|
---------------------------------------------------------------
|_UUId0231|map[session2]{"name":123},....repeat |MemorySession|
---------------------------------------------------------------
*/
type MemorySessionMgr struct {
	sessonMgr map[string]Session
	rwlock  sync.RWMutex
}

func NewMemorySessionMgr()*MemorySessionMgr{
	return &MemorySessionMgr{
			sessonMgr: make(map[string]Session,1024),
	}
}

func(m *MemorySessionMgr)Init(addr string,options...string)(err error)  {
	return
}
func(m *MemorySessionMgr)CreateSession()(session Session,err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	//go get github.com/satori/go.uuid

	id := uuid.NewV4()
	fmt.Printf("%v",id)
	sessionId = id.String()
	session= NewMemorySession(sessionId)
	m.sessonMgr[sessionId]=session
	return
}
func(m *MemorySessionMgr)Get(sessionId string)(session Session,err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session,ok:=m.sessonMgr[sessionId]
	if !ok{
		err = errors.New("get session failed ,sessionId or undefined")
		return
	}
	return
}





