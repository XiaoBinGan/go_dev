package main

import (
	"fmt"
	"sort"
)
/*
	通过切割数组获得切片
*/
func slice1()  {
		//1.切片的定义
		var a1 []int           //定义一个存放int类型元素的切片
		var b1 []string        //定义一个存放string类型元素的切片
		fmt.Println(a1,b1)     //切片唯一合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0但是我们不能说一个长度和容量都是0的切片一定是nil
								// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
								// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
								// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
								// 所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。
		fmt.Println(a1 == nil) //true 因为指向一个空数组没有开辟内存空间所以是nil
		fmt.Println(b1 == nil) //true
	
		//1.1初始化切片
		a1 = []int{1,2,3} 
		b1 = []string{"黄浦","浦东","静安"}
		fmt.Println(a1,b1)     //初始化了底层数组但是不需要指定长度因为切片会自己扩容
		fmt.Println(a1 == nil) //false 初始化了底层数组开辟内存空间所以是不等于nil
		fmt.Println(b1 == nil) //false
		//1.2.基本长度和容量
		fmt.Printf("a1 :len(a1):%d cap(a1) :%d\n", len(a1),cap(a1))
		fmt.Printf("b2 :len(b1):%d cap(a1) :%d\n", len(b1),cap(b1))
	
		//2.由数组得到切片
		a2 :=[...]int{1,3,5,7,9,11,13}
		a3 :=a2[0:4]	        //基于数组切割得到的切片,包含左边的位图不包含右边的(左闭右开)
		fmt.Println("a3:",a3)   //[1 3 5 7]
		fmt.Printf("a3:%T\n",a3)//[]int
		a4 := a2[1:6]			//[3 5 7 9 11]
		fmt.Println("a4:",a4)
		a5 :=a2[:4]				//[0 :4] [1 3 5 7]
		a6 :=a2[2:]			    //[2:len(a2)] [5 7 9 11 13]
		a7 :=a2[:]				//[0:len(a2)]  [1 3 5 7 9 11 13]
		fmt.Printf("a5:%v ,a6:%v ,a7:%v \n",a5,a6,a7)
		//这里切片的容量等同于底层数组的容量
		fmt.Printf("a5 :len(a5):%d cap(a5) :%d\n", len(a5),cap(a5))
		//这里切片的长度等于切割的数组返回的长度
		//容量等于切片的第一个元素到最后的元素
		fmt.Printf("a5 :len(a6):%d cap(a6) :%d\n", len(a6),cap(a6))
		a8 :=a6[4:] 			//再切割 [13]
		fmt.Printf("a8 :len(a8):%d cap(a8) :%d\n", len(a8),cap(a8))
		//切片是引用类型都指向底层的一个数组
		fmt.Println("a6:",a6)
		a2[6] = 2400 			//修改底层数组的值  后面的切片都会改变
		fmt.Println("a6:",a6)
		fmt.Println("a8:",a8)
}
/*
	make函数创造切片
	切片就是一个框,框柱了一块连续的内存
	切片属于引用类型,真正的数据都是保存在底层数组里的
*/
func makeSlice()  {
	/*make([]T, size, cap)
		T:切片的元素类型
		size:切片中元素的数量
		cap:切片的容量
		如果cap不写则默认容量和长度是一致的
	*/
	 s1 :=make([]int,5,10)
	 fmt.Printf("s1:%v ,len(%d),cap(%d) \n",s1,len(s1),cap(s1))
	 s2 :=make([]int,0,10)
	 fmt.Printf("s2:%v ,len(%d),cap(%d) \n",s1,len(s2),cap(s2))
	 //切片就是赋值
	 s3 :=[]int{1,3,5}
	 s4 :=s3	//这时两个参数指向同一个底层数组
	 fmt.Println(s3,s4)
	 s3[1]=100
	 fmt.Println(s3,s4)
	//slice 的遍历
	 s5 :=[]int{1,3,5}
	//1.for循环遍历
	 for i := 0; i < len(s5); i++ {
		 fmt.Println(i,s5[i])
	 }
	//2.for  range 
	for i,v := range s5 {
		fmt.Println(i,v)
	}

} 
/*
	append()为切片追加元素
*/
func appendSlice()  {
	s1 :=[]string{"上海","浦东","张江"}
	fmt.Printf("s1:%v ,len(%d),cap(%d) \n", s1,len(s1),cap(s1)) //扩容之前打印一下切片的长度和容量
																/*错误示范
																	s1[3] = "金桥"	  //这样的操作会报错因为数组越界了
																	fmt.Println(s1)//panic: runtime error: index out of range [3] with length 3
																*/
	s1 = append(s1,"金桥")										 //append()追加元素,原来的底层数组放不下的时候,Go底层会把底层数组换一个
	fmt.Println("s1:",s1)										//调用append函数必须 用变量接收返回值(注意点,最好使用原来的参数接收,新参数也可以但是不推荐)
	fmt.Printf("s1:%v ,len(%d),cap(%d) \n", s1,len(s1),cap(s1)) //扩容之后打印一下切片的长度和容量
	s2 :=append(s1,"金桥路")									 //测试这里相当于新开的内存地址和原来的不是一个底层数组
	fmt.Printf("s2:%v,\ns1:%v\n",s2,s1)							//[上海 浦东 张江 金桥 金桥路]   [上海 浦东 张江 金桥]
	fmt.Printf("s2:%v ,len(%d),cap(%d)\n", s2,len(s2),cap(s2))
	s1 =append(s1,"金桥","金桥","金桥")
	fmt.Printf("s1:%v ,len(%d),cap(%d) \n", s1,len(s1),cap(s1)) //扩容之后打印一下切片的长度和容量 容量没超过还是原来的容量,不是很复杂的超了直接容量翻倍
	s3 :=[]string{"张江","漕河泾","闸北"}
	s1 =append(s1,s3...)
	fmt.Printf("s1:%v ,len(%d),cap(%d) \n", s1,len(s1),cap(s1)) // 容量没超过还是原来的容量,不是很复杂的超了直接容量翻倍



	/*
			首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。

			否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），

			否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
			即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）

			如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

			需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	*/
}
/*
	copySlice
*/
func copySlice()  {
	a1 :=[]int{1,3,5}
	a2 :=a1						//赋值
	var a3 = make([]int,3,3)	//创建内存空间 长度为3容量为3
	copy(a3,a1)					//如果a3容量接受不下a1的内容那么切片会自动截断后面的内容舍弃    
	fmt.Println(a1,a2,a3)
	a1[0] = 100					//a3相当于拷贝一份a 1的内容所以a1的内容修改了a3并不会受到影响
	fmt.Println(a1,a2,a3)		


	x1 :=[...]int{1,3,5,7,9,12}		//数组
	s1 :=x1[:]				    //切片获取第一位到最后一位的热凝
	fmt.Printf("s1:%v s,len(%d),cap(%d) \n",s1,len(s1),cap(s1))
	//切片不存具体的值 切片对应一个底层数组 底层数组都是占用的一块连续的内存
	s1 = append(s1[:1],s1[2:3]...)//优先保证切割的内容进行存储如果没有存放满就自右向左缺几位填入几位
	
	s1[0] = 100
	fmt.Println(x1)
	fmt.Println(s1)
}
/*
 exercises
*/
func exercisesSlice()  {
	var a = make([]int, 5, 10)
	fmt.Println("a:",a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)




	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])//对切片进行排序
	fmt.Println("a1:",a1)

}



func main()  {
	// slice1()       //	通过切割数组获得切片
	// makeSlice()    //	make函数创造切片
	// appendSlice()  //	append()为切片追加元素
	copySlice()
	exercisesSlice()
}
