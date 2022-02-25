package main

import (
	"os"
	"fmt"
)



/*
	函数学生管理系统
	写一个系统查看、新增学生、删除学生
*/

type student struct{
	id int64
	name string
}

type studentManeger struct{
	allstudent map[int64]*student
}


//show all student
func (s *studentManeger)showAllstudent(){
	for k, v := range s.allstudent {//for range all student 
		fmt.Printf("学号%d,姓名%s\n", k,v.name)
	}
}
func (s *studentManeger)addStudent(){
	//add the student to allstudent
	var (
		id int64
		name string
	)
	fmt.Println("学生新增")
	//assign user Scanln value to id
	fmt.Print("请输学生学号:")
	fmt.Scan(&id)
	//Judge whether the students are present or not
	it,ok :=s.allstudent[id]
	if ok {
		fmt.Println("学生已存在请重新录入",it)
		return
	}
	//assign user Scanln value to name
	fmt.Print("请输学生姓名:")
	fmt.Scan(&name)
	newStu :=&student{
		id:id,
		name:name,
	}
	s.allstudent[id]=newStu
}
func  (s *studentManeger)deleteStudent(){
	var id int64
	fmt.Print("请输入需要删除学生的id:")
	fmt.Scanln(&id)
	it,ok :=s.allstudent[id]//这里的it拿到的是allstudent中单个student结构体
	if !ok {
		fmt.Println("学生不存在")
		return
	}
	delete(s.allstudent,id)
	fmt.Printf("删除的学生信息如下%v",it)
}



func main() {
	s :=&studentManeger{
		allstudent:make(map[int64]*student,88),
	}
	for{
		//1.first print menu
		fmt.Println("welcome to the student system")
		fmt.Println(`
			please input your choose number
			1.cat all student
			2.add student
			3.delete student
			4.exit 
		`)
		fmt.Print("what are you want to do ?")
		//2.second waiting user input ,than to do .
		var choice int //get user input for choice
		fmt.Scanln(&choice)//Assign the input to choice
		fmt.Printf("你选择了%d这个选项:",choice)
		switch choice {
		case 1:
			s.showAllstudent()
		case 2:
			s.addStudent()
		case 3:
			s.deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("get out")			
		}
	}
}