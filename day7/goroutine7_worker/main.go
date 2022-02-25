package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)
// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
// 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
// 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 主goroutine从resultChan取出结果并打印到终端输出
type jobs struct{//接收随机的int64的数
	value int64
}
type result struct{//接收jobs随机数的结构体
	job *jobs
	sum int64
}
var jobChan =make(chan *jobs, 50)  //存储jons结构体
var resChan =make(chan *result, 50)//存储结果结构体
var wg sync.WaitGroup			   //创建程序退出条件
func doit(j chan <- *jobs){		   //循环生成int64类型的随机数，发送到jobChan	
	defer wg.Done()				   //执行结束-1
	for {						   //循环的去放随机数
		x :=rand.Int63()		   //循环的去生成int64随机数
		newJob :=&jobs{
			value:x,
		}
		time.Sleep(time.Millisecond*500)
		j<-newJob
	}
}
func getsum(j <-chan *jobs,r chan<-*result)  {//从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	defer wg.Done()							  //执行结束-1
	for{									  //死循环的去读取
		job :=<-j							  //去除joschan当中的job结构体进行操作
		n := job.value						  //去除int64随机数
		sum :=int64(0)						  //因为结构体存储的事int64的数 所以这里申明的时候转成int64
		for n >0{							  //只要n大于0就继续循环添加	
			sum+=n%10						  //每次取余获取到最后一位
			n =n/10							  //int64计算是不会出现浮点数的所以这里都是整数  得出的结构继续循环直到变成0   0.213123123/10=0
		}
		newr :=&result{						  //将得出的结果和job随机数存入result结构体
			job:job,
			sum:sum,
		}
		r <-newr							 //将生成的result存入reschan
	}
}


func main() {
	wg.Add(1)								//需要执行多少个goroutine就add几位数
	go doit(jobChan)						//生成随机数
	for i := 0; i < 24; i++ {				//开启24个任务线程
		go getsum(jobChan,resChan)			//从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	}

	for v:= range resChan {					//循环的去读取reschan的内容
		fmt.Printf("value:%d,sum:%d\n",*v.job,v.sum)
	}
	wg.Wait()								//线程全部结束之后结束函数执行
}
// n :=int64(32)
// n=n/10//int类型只能算出来整数
// fmt.Println(n)



// var b float64
// n :=int64(32)
// b=float64(n/10) //int类型只能算出来整数
// fmt.Println(b)3


// sum :=int64(0)
// n :=int64(12312)
// for n>0{
// 	sum+=n%10
// 	fmt.Println(n)
// 	n=n/10
// 	fmt.Println(sum)
// 	time.Sleep(time.Second)
// }