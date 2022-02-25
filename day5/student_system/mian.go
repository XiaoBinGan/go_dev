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
//student constructe function 
func newStudent(id int64,name string) *student {
		return &student{
			id:id,
			name:name,
		}
	
}

var (
	allstudent map[int64]*student
)


//show all student
func showAllstudent(){
	for k, v := range allstudent {//for range all student 
		fmt.Printf("学号%d,姓名%s\n", k,v.name)
	}
}
func addStudent(){
	//add the student to allstudent
	var (
		id int64
		name string
	)
	fmt.Println("学生新增")
	//assign user Scanln value to id
	fmt.Print("请输学生学号:")
	fmt.Scan(&id)
	//assign user Scanln value to name
	fmt.Print("请输学生姓名:")
	fmt.Scan(&name)
	newStu :=newStudent(id,name)
	allstudent[id]=newStu
}
func deleteStudent(){
	var id int64
	fmt.Print("请输入需要删除学生的id:")
	fmt.Scanln(&id)
	delete(allstudent,id)
}



func main() {
	allstudent = make(map[int64]*student,88)
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
			showAllstudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("get out")			
		}
	}
}