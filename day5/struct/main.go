package main

import (
	"fmt"
)
//结构体是值类型 
// struct的内存布局：struct中的所有字段在内存是连续的：
type student struct{
	Name string
	Age  int
	score int 
}
/*初始化方法1*/
func structDemo()  {
	var stu1 student
	stu1.Name="xiu"
	stu1.Age=18
	stu1.score=99
	fmt.Printf("Name:%p\n",&stu1.Name)
	fmt.Printf("Name:%p\n",&stu1.Age)
	fmt.Printf("Name:%p\n",&stu1.score)
}

/*初始化方法2*/
func structDemo1()  {
	var stu2 *student = &student{
		Name:"xiu",
		Age:28,
		score:199,
	}
	fmt.Printf("Name:%p\n",&stu2.Name)
	fmt.Printf("Name:%p\n",&stu2.Age)
	fmt.Printf("Name:%p\n",&stu2.score)
}
/*初始化方法3*/
func structDemo2()  {
	var stu2 = student{
		Name:"xiu",
		Age:28,
		score:199,
	}
	fmt.Printf("Name:%p\n",&stu2.Name)
	fmt.Printf("Name:%p\n",&stu2.Age)
	fmt.Printf("Name:%p\n",&stu2.score)
}
/*初始化方法4*/
func structDemo3()  {
	var stu2 = &student{
		"xiu", //使用值列表的形式初始化,值的顺序要和结构体定义时的字段顺序一致
		28,	   //使用值列表的形式必须填满值缺一个都会报错
		199,
	}
	fmt.Printf("%#v\n",stu2)
	fmt.Printf("stu2:%p\n",stu2)
	fmt.Printf("&stu2:%p\n",&stu2)
	f1 :=func (x *student)  {//根据内存地址找到那个变量,修改的就是原来的变量
		x.Age=19
	}
	f1(stu2)
	fmt.Println(stu2.Age)
}
/*go语言中函数参数永远是拷贝*/
func structDemo4()  {
	 f :=func (x student)  {//go语言中函数参数永远是拷贝
		x.Age=19
	}
	 f1 :=func (x *student)  {//根据内存地址找到那个变量,修改的就是原来的变量
		x.Age=19
	}
	var p student
	p.Name="wangxin"
	p.Age=20
	f(p)
	f1(&p)
	fmt.Println(p.Age)


	var p1 = new(student)   //这里拿到的就是一个内存地址赋值给了p1
	p1.Age=29
	fmt.Printf("%T\n",p1)	//这里的类型就是main.student
	fmt.Printf("%p\n", p1)  //p1保存的值就是一个内存地址
	fmt.Printf("%p\n", &p1)	//&p1拿的是p1的内存地址
	

}

func main() {
	structDemo()
	structDemo1()
	structDemo2()
	structDemo3()
	structDemo4()

}


/*
用来自定义复杂数据结构
2. struct里面可以包含多个字段（属性）
3. struct类型可以定义方法，注意和函数的区分
4. struct类型是值类型
5. struct类型可以嵌套
6. Go语言没有class类型，只有struct类型
1. struct 声明：
				type 标识符 struct {
					field1 type
					field2 type
				}
例子
				type Student struct {
					Name string
					Age int
					Score int
				}
2. struct 中字段访问：和其他语言一样，使用点
例子
				var stu Student

				stu.Name = “tony”
				stu.Age = 18
				stu.Score=20

				fmt.Printf(“name=%s age=%d score=%d”, 
					stu.Name, stu.Age, stu.Score)
3.  struct定义的三种形式：
				var stu Student
				var stu *Student = new (Student)
				var stu *Student = &Student{}
1）其中b和c返回的都是指向结构体的指针，访问形式如下：
a. stu.Name、stu.Age和stu.Score或者 (*stu).Name、(*stu).Age等
4. struct的内存布局：struct中的所有字段在内存是连续的：

*/