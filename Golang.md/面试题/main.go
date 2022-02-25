package main

var N =10

func main() {
//面试题
}


/*
1.golang 语言初始化的顺序
	main包
	import
	全局 const
	全局 var
	init
	main 函数


2.func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu :=&sync.Mutex{}
	wg.Add(N)
	for i:=0;i<N; i++ {
		go func(i int) { 将i通过参数形式传递进入匿名函数执行 不然
			defer wg.Done()
			mu.Lock()
			m[i]=i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	fmt.Printf("%#v",m)
	//for _, v := range m {
	//	fmt.Printf("%#v\n",v)
	//}
}

3.func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu :=&sync.Mutex{}
	wg.Add(N)
	for i:=0;i<N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock() 并发问题 加上锁就不会出现
 			m[rand.Int()]=rand.Int()
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
}
4。func main() {
	type S struct {
		name string
	}
    题目m :=map[string]S{"x":S{"one"}}
	//map里的结构体无法直接寻址，必须取址
  	m :=map[string]*S{"x":&S{"one"}}
  	m["x"].name="two"
  	println(m["m"].name)
}
5。func main() {
	type Result struct {
		Status int
		//status int 问题就出现在这里的字段是小写 解码的时候会外面会没法使用
	}
	var data =[]byte(`{"status":200}`)
	result :=&Result{}
	if err :=json.Unmarshal(data,&result);err!=nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Printf("result=%+v",result)

}

6。描述golang中的stack和heap的区别，分别在什么情况下会分配stack？又会在何时会分配到heap中？
区别：
	栈（stack）：由编译器自动分配和释放。存放变量名，各种名
	堆：在C里面由程序员分配和释放内存，go自动了，存栈变量的数据


make(xxx) a:=3 a就是在栈 3就是在堆


*/