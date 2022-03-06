package main

import "fmt"


//new(T) 返回 T 的指针 *T 并指向 T 的零值。
//make(T) 返回的初始化的 T，只能用于 slice，map，channel。



// 下面的面的代码是等价的，new(int) 将分配的空间初始化为 int 的零值，也就是 0，
//并返回 int 的指针，这和直接声明指针并初始化的效果是相同的。
func testNew()  {
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1) //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0

	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2) //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0
}

//make 只能用于 slice,map,channel
//make 只能用于 slice，map，channel 三种类型，make(T, args) 返回的是初始化之后的 T 类型的值，
//这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。
func testMake()  {
	var s1 []int
	if s1==nil{
		fmt.Printf("s1 is nil --> %#v \n",s1)//[]int(nil)
	}
	s2 :=make([]int,3)
	if s2==nil{
		fmt.Printf("s2 is nil l --> %#v \n",s2)//[]int(nil)
	}else{
		fmt.Printf("s2 is not nil --> %#v \n",s2)//[]int{0,0,0}
	}
}
//slice 的零值是nil, 使用make之后的slice是一个初始化的slice，
//即slice的长度、容量、底层指向的array都被make完成初始化。
//此时的slice内容类型 int的零值填充形式是【0 0 0】，map和channel也是类似的

func testMakeMapChan()  {
    //map
	var m1 map[int]string
	if m1 ==nil{
		fmt.Printf("m1 is nil --> %#v \n",m1)// map[int]string(nil)
	}
	m2 :=make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
	}
	//channel
	var  c1 chan string
	if c1 ==nil{
		fmt.Printf("c1 is nil --> %#v \n",c1)//(chan string)(nil)
	}
	c2 :=make(chan string,2)
	if c2 ==nil{
		fmt.Printf("c2 is nil --> %#v \n",c2)
	}else {
		fmt.Printf("c2 is not nil -->%#v \n",c2)//(chan string)(0xc0000561e0)
	}
}
/**
    很重要的点
    如果不是特殊声明，go的函数默认都死按值传参，即通过函数传递的参数是值得副本，在函数内部对值得修改不影响值得本身
    但是但是但是
    make(T,args)返回得值通过函数传递参数到函数内部可以直接修改，
    即map、channel、slice通过函数参数传递之后在函数的内部修改将影响函数外部的值
    这说明make(T，args)返回的不是引用类型，golang函数调用总是传值。 可以更改生效是因为两个slice指向同一个地址空间，
    在函数内部可以直接更改原始值，对map和channel也是如此
 */
func modifySlice(s []int){
	s[0]=1
}
func modifyMap(m map[int]string)  {
	m[10]="string"
}
func modifyChan(c chan string)  {
	c <-"string"
}
func testMakeFuncProps()  {
	//1、Slice
	s2:=make([]int,3)
	fmt.Printf("%#v \n",s2)//[]int{0, 0, 0}
	modifySlice(s2)
	fmt.Printf("%#v \n",s2)//[]int{1, 0, 0}
	//2、Map
	m2:=make(map[int]string)
	if m2==nil{
		fmt.Printf("m2 is nil-->%#v \n",m2)
	}else{
		fmt.Printf("m2 is not nil--> %#v \n",m2)
	}
	modifyMap(m2)
	fmt.Printf("m2 is changed-->%#v \n",m2)
	//3、channel
	c2 :=make(chan string)
	if c2 ==nil{
		fmt.Printf("c2 is nil --> %#v \n",c2)
	}else{
		fmt.Printf("c2 is not  nil --> %#v \n",c2)
	}
	go modifyChan(c2)
	//您的chan是未缓冲的，所以当您试图通过该通道发送一个值时，它会永远阻塞，等待别人获取一个值。
	//您需要启动一个新的goroutine，或使通道缓冲并使用它作为一个数组。
	//如果make了内存就不需要goroutine
	fmt.Printf("c2 is changed-->%#v \n",c2 )
}



func testNewStruct()  {
	type Student struct {
		Name string
		Age  int
	}
	//1、声明初始化
	var st1 Student
	st1 =Student{
		Age:18,
	}
	st1.Name="张三"//如果这条语句在上面语句之前 会被st1=Student{}覆盖
	fmt.Printf("st1 -->%#v \n",st1)
    //2、struct literal 字面量初始化
    st2 :=Student{}
    fmt.Printf("st2 ---->%#v \n",st2)
    st2.Age=12
	fmt.Printf("st2 ---->%#v \n",st2)
    //3、指针初始化
    st3 :=&Student{}
	fmt.Printf("st3 ---->%#v \n",st3)
    st3.Name="St3"
	fmt.Printf("st3 ---->%#v \n",st3)
    //4、new 关键字初始化
    st4 :=new(Student)
    fmt.Printf("st4---->%#v \n",st4)
    st4.Name="st4"
    fmt.Printf("st4--->%#v \n",st4)
    //5.声明指针并用new关键字初始化
    var st5 *Student = new(Student)
    fmt.Printf("st5 %#v \n",st5)
    st5.Age=19
    st5.Name="st5"
    fmt.Printf("st5---->%#v \n",st5)
    //st1和st2是相同的类型，都是student类型的值，st1通过var声明，student的filed 自动初始化类型零值，
    //st2通过字面量的完成初始化. st3~st4~st5的类型是一样的，都是Student类型，都是student的指针student.但是所有的student都可以直接被filed，
    //如果X是可寻地址，&x的fild集合包含吗m.x.m和(&x)是等同的，go自动转换，也就说st1和st2调用是等价的，go在下面做转换。
    //因此可以直接使用struct literal的方式创建对象，能到和new创建的一样的效果而不需要使用new关键字
}


func main()  {
	testNew()
	testMake()
	testMakeMapChan()
	testMakeFuncProps()
	testNewStruct()
}