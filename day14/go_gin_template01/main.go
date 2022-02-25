package main

import (
	"fmt"
	"html/template"
	"net/http"
)



func hd(w http.ResponseWriter,r *http.Request) {
	t,err :=template.ParseFiles("./index.tpl")
	if err!=nil{
		fmt.Printf("read template failed,err:%#v\n",err)
		return
	}
	t.Execute(w,"球球")

}


func main() {
	http.HandleFunc("/",hd)
	http.ListenAndServe(":9000",nil)
}
