package session
//定义一个管理者管理所以的session
/**
------------------MemorySessionMgr-----------------------------
|sessionId|map[string]interface{}				| 备注 		  |
---------------------------------------------------------------
|_UUId0xk1|map[session1]{"name":123},....repeat |MemorySession|
---------------------------------------------------------------
|_UUId0231|map[session2]{"name":123},....repeat |MemorySession|
---------------------------------------------------------------
*/
type SessionMgr interface {
	Init(addr string,options... string)(err error)
	CreateSession()(Session Session,err error)
	Get(sessionId string)(session Session,err error)
}
