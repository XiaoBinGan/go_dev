package main

import (
	"net"
	"fmt"
)
func process(conn net.Conn)  {
    defer conn.Close() //close connection
    for {
        buf :=make([]byte,1000)//创建一个byte数字
        n,err:=conn.Read(buf)//使用net.Conn.Read读取这个字节数组
        if err!=nil {//容错
            fmt.Println("READ failed,err:",err)
        }
        fmt.Printf("%v\n",string(buf[0:n]))
        //tosring 这个字符数组从0位到最后一位
    }
}
func main() {
	fmt.Println("start server ...")
	listen,err :=net.Listen("tcp","0.0.0.0:50000")
	if err!=nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	for {
		Conn,err :=listen.Accept()
		if err!=nil {
			fmt.Println("Accept is failed,err:",err)
		}
		go process(Conn)
	}
}


/*
func Listen
func Listen(network, address string) (Listener, error)
Listen announces on the local network address.

The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".

For TCP networks, if the host in the address parameter is empty or a literal unspecified IP address, Listen listens on all available unicast and anycast IP addresses of the local system. To only use IPv4, use network "tcp4". The address can use a host name, but this is not recommended, because it will create a listener for at most one of the host's IP addresses. If the port in the address parameter is empty or "0", as in "127.0.0.1:" or "[::1]:0", a port number is automatically chosen. The Addr method of Listener can be used to discover the chosen port.

See func Dial for a description of the network and address parameters.




type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)

    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error

    // Addr returns the listener's network address.
    Addr() Addr
}


Conn is a generic stream-oriented network connection.

Multiple goroutines may invoke methods on a Conn simultaneously.

type Conn interface {
    // Read reads data from the connection.
    // Read can be made to time out and return an Error with Timeout() == true
    // after a fixed time limit; see SetDeadline and SetReadDeadline.
    Read(b []byte) (n int, err error)

    // Write writes data to the connection.
    // Write can be made to time out and return an Error with Timeout() == true
    // after a fixed time limit; see SetDeadline and SetWriteDeadline.
    Write(b []byte) (n int, err error)

    // Close closes the connection.
    // Any blocked Read or Write operations will be unblocked and return errors.
    Close() error

    // LocalAddr returns the local network address.
    LocalAddr() Addr

    // RemoteAddr returns the remote network address.
    RemoteAddr() Addr

    // SetDeadline sets the read and write deadlines associated
    // with the connection. It is equivalent to calling both
    // SetReadDeadline and SetWriteDeadline.
    //
    // A deadline is an absolute time after which I/O operations
    // fail with a timeout (see type Error) instead of
    // blocking. The deadline applies to all future and pending
    // I/O, not just the immediately following call to Read or
    // Write. After a deadline has been exceeded, the connection
    // can be refreshed by setting a deadline in the future.
    //
    // An idle timeout can be implemented by repeatedly extending
    // the deadline after successful Read or Write calls.
    //
    // A zero value for t means I/O operations will not time out.
    //
    // Note that if a TCP connection has keep-alive turned on,
    // which is the default unless overridden by Dialer.KeepAlive
    // or ListenConfig.KeepAlive, then a keep-alive failure may
    // also return a timeout error. On Unix systems a keep-alive
    // failure on I/O can be detected using
    // errors.Is(err, syscall.ETIMEDOUT).
    SetDeadline(t time.Time) error

    // SetReadDeadline sets the deadline for future Read calls
    // and any currently-blocked Read call.
    // A zero value for t means Read will not time out.
    SetReadDeadline(t time.Time) error

    // SetWriteDeadline sets the deadline for future Write calls
    // and any currently-blocked Write call.
    // Even if write times out, it may return n > 0, indicating that
    // some of the data was successfully written.
    // A zero value for t means Write will not time out.
    SetWriteDeadline(t time.Time) error
}




*/