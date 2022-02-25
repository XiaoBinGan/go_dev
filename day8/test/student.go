package main

import (
	"io/ioutil"
	"encoding/json"
)
//Student struct
type Student struct {
	Name string 
	Age int
	score int
}
/*
*Save 
*Params *Student
*/
func (p *Student)Save()(err error){
	data,err :=json.Marshal(p)
	if err !=nil {
		return
	}
	err = ioutil.WriteFile("./stu.dat",data,0755)
	return
}
func (p *Student)Load()(err error){
	data,err :=ioutil.ReadFile("./stu.dat")
	if err !=nil {
		return
	}
	err = json.Unmarshal(data,p)
	return
} 
