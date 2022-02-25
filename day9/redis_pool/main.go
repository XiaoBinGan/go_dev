package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		MaxIdle: 20,
		MaxActive: 0,
		IdleTimeout:300,
		Dial:func ()(redis.Conn,error){
			return redis.Dial("tcp", "localhost:6379"), nil
		},
	}
}
func main() {
		// 	Get gets a connection. The application must close the returned connection.
		// This method always returns a valid connection so that applications can defer
		// error handling to the first use of the connection. If there is an error
		// getting an underlying connection, then the connection Err, Do, Send, Flush
		// and Receive methods return that error.
	c :=pool.Get()  //pool实现了Get方法 返回一个conn 
	defer c.Close() //适应结束关闭连接

	//接受用户的指令
	_,err :=c.Do("Set","abc",100)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//int是一种帮助器，它将命令应答转换为整数。如果我犯错 等于nil，然后Int返回0,err。否则，Int将转换
	r,err :=redis.Int(c.Do("Get","abc"))
	if err!=nil {
		fmt.Println("get value failed,err:",err)
		return
	}
	fmt.Println(r)
	pool.Close()



}






// Pool maintains a pool of connections. The application calls the Get method
// to get a connection from the pool and the connection's Close method to
// return the connection's resources to the pool.
//
// The following example shows how to use a pool in a web application. The
// application creates a pool at application startup and makes it available to
// request handlers using a package level variable. The pool configuration used
// here is an example, not a recommendation.
//
//  func newPool(addr string) *redis.Pool {
//    return &redis.Pool{
//      MaxIdle: 3,
//      IdleTimeout: 240 * time.Second,
//      Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
//    }
//  }
//
//  var (
//    pool *redis.Pool
//    redisServer = flag.String("redisServer", ":6379", "")
//  )
//
//  func main() {
//    flag.Parse()
//    pool = newPool(*redisServer)
//    ...
//  }
//
// A request handler gets a connection from the pool and closes the connection
// when the handler is done:
//
//  func serveHome(w http.ResponseWriter, r *http.Request) {
//      conn := pool.Get()
//      defer conn.Close()
//      ...
//  }
//
// Use the Dial function to authenticate connections with the AUTH command or
// select a database with the SELECT command:
//
//  pool := &redis.Pool{
//    // Other pool configuration not shown in this example.
//    Dial: func () (redis.Conn, error) {
//      c, err := redis.Dial("tcp", server)
//      if err != nil {
//        return nil, err
//      }
//      if _, err := c.Do("AUTH", password); err != nil {
//        c.Close()
//        return nil, err
//      }
//      if _, err := c.Do("SELECT", db); err != nil {
//        c.Close()
//        return nil, err
//      }
//      return c, nil
//    },
//  }
//
// Use the TestOnBorrow function to check the health of an idle connection
// before the connection is returned to the application. This example PINGs
// connections that have been idle more than a minute:
//
//  pool := &redis.Pool{
//    // Other pool configuration not shown in this example.
//    TestOnBorrow: func(c redis.Conn, t time.Time) error {
//      if time.Since(t) < time.Minute {
//        return nil
//      }
//      _, err := c.Do("PING")
//      return err
//    },
//  }
//