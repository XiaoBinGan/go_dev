package main

import (
	"path"
	"time"
	"fmt"
	"os"
)

func main() {
	//1.打开文件 获取os.file的 指针对象
	fileObj,err :=os.Open("./mian.go")
	if err!=nil{
		fmt.Println("open file failed ,err:",err)
		return
	}
	//2.通过Stat方法获取到fileInfo 接口对象
	fileInfo,e :=fileObj.Stat()
	if e!=nil{
		fmt.Println("get file info failed ,err:",e)
		return
	}
	//3.获取到文件的具体大小
	fmt.Printf("%db\n",fileInfo.Size())
	fmt.Printf("%s\n",fileInfo.Name())//不包含全路径的文件名
	now :=time.Now().Format("20060102150405")
	odlName :=fileInfo.Name()
	str :=path.Join("%s.bak%s",odlName,now)
	fmt.Println(str)
	
}




// //// A FileInfo describes a file and is returned by Stat and Lstat.
// type FileInfo interface {
// 	Name() string       // base name of the file
// 	Size() int64        // length in bytes for regular files; system-dependent for others
// 	Mode() FileMode     // file mode bits
// 	ModTime() time.Time // modification time
// 	IsDir() bool        // abbreviation for Mode().IsDir()
// 	Sys() interface{}   // underlying data source (can return nil)
// }