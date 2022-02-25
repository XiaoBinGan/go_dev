package main

import (
	"time"
	"sync"
	"fmt"
)
 var (
	m = make(map[int]uint64)
	lock sync.Mutex   //互斥锁
 )
type task struct{
	n int 
}
/*
1.1000以内的阶乘
2.for循环0_1000

*/

func calc(t *task){	
	var sum uint64 
	sum =1
	for i := 1; i < t.n; i++ {
		sum *=uint64(i)
	}
	lock.Lock()
	m[t.n]=sum
	lock.Unlock()
}
func jiechen()  {
  for i := 0; i < 1000; i++ {
	t :=&task{n:i} 
	go calc(t)
  }
  time.Sleep(time.Second * 10)
  lock.Lock()
  for k,v := range m {
	  fmt.Printf("%d! = %v\n",k,v)
  }
  lock.Unlock()
}


/*
  使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，
  其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，
  多个goroutine同时等待一个锁时，唤醒的策略是随机的。
*/
var x=0  //这里不
var wg sync.WaitGroup
var locks sync.Mutex
func sMutex()  {
	fn :=func(){
		for i := 0; i < 5000; i++ {
			locks.Lock()	
			x++						//两个函数本应该交替去+1但是因为goroutine的关系的所以寸会存在数据的竞争导致结果不符合
			locks.Unlock()
		}
		wg.Done()
	}
	wg.Add(2)
	go fn()
	go fn()
	wg.Wait()
	fmt.Println(x)
}




func main() {
	// jiechen()
	sMutex()
}
