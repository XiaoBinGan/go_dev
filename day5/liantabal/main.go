package main

import (
	"math/rand"
	"fmt"
)


type Student struct{
	Name   string
	Age    int 
	Score  float32
	next    *Student 
}
/*
  遍历链表所以子节点
*/
func forr(a *Student)  {
	var z *Student =a
	for z!=nil {
		fmt.Println(*z)
		z=z.next
	}
}
/*
	循环添加子列表
*/
func forkids(a *Student)  {
	c :=a							   
	for i := 0; i < 10; i++ {
		b :=Student{
			Name:fmt.Sprint("p:%d",i), /*按照格式输出一个符合格式的string*/
			Age:rand.Intn(100),        /*返回一个0到100之间的随机数*/
			Score:rand.Float32()*100,  /*返回一个0.0到1.0之间的值然后乘以100得出0到100指甲的随机数*/
		}
		c.next=&b
		c=&b
		// fmt.Println(c)
	}
}
/*
  demo1
*/
func demo1()  {
	var a =Student{
		Name:"qwe",
		Age:19,
		Score:120,
	}
	var b =Student{
		Name:"asd",
		Age:19,
		Score:10,
	}
	a.next=&b
	// fmt.Print(a)
	forr(&a)

}
/*
  demo2
  链表 尾部插入法
*/
func demo2()  {
	var a =Student{
		Name:"x",
		Age:20,
		Score:99,
	}
	// var a *Student =&Student{
	// 	Name:"x",
	// 	Age:20,
	// 	Score:99,
	// }
	forkids(&a)
	forr(&a)
}


/*
  demo3
  链表的头部插入法
  		9->8->7->6->5->4->3->2->1->0
		a是链表的最底层
		通过a不断将父元素复制给原来承载底层变量然后断的往上赋值直到最顶层
*/
func demo3()  {
	    var a *Student=new(Student)
		a.Name="xiu"
		a.Age=19
		a.Score=99

	for i := 0; i < 10; i++ {
		b :=Student{
			Name:fmt.Sprintf("b%d",i),
			Age:rand.Intn(100),
			Score:rand.Float32()*100,
		}
		b.next=a
		a=&b
	}
	forr(a)
	
 }

func main() {
	// demo1()
	demo2()
	demo3()
}   

/*
	5. 链表定义
	type Student struct {
		Name string
		Next* Student
	}
	每个节点包含下一个节点的地址，这样把所有的节点串起来了，通常把
	链表中的第一个节点叫做链表头

*/



/*
  func demo2()  {
		var a =Student{
			Name:"x",
			Age:20,
			Score:99,
		}
	 	// var a *Student =&Student{
	 	// 	Name:"x",
	 	// 	Age:20,
	 	// 	Score:99,
	 	// }

		
		tail :=&a
		for i := 0; i < 10; i++ {
				b :=Student{//申明的时候直接初始化
					Name:fmt.Sprint("b%d",i),
					Age:rand.Intn(100),
					Score:rand.Float32()*100,
				}
				tail.next=&b
				tail=&b
		}
		forr(a)
	
	}
*/