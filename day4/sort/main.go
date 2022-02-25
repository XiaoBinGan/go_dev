package main
/**
排序操作主要都在 sort包中，导入就可以使用了
import(“sort”)
sort.Ints对整数进行排序， 
sort.Strings对字符串进行排序, 
sort.Float64s对浮点数进行排序.
sort.SearchInts(a []int, b int) 从数组a中查找b，前提是a必须有序
sort.SearchFloats(a []float64, b float64) 从数组a中查找b，前提是a必须有序
sort.SearchStrings(a []string, b string) 从数组a中查找b，前提是a必须有序
*/
import (
	"sort"//排序内置包
	"fmt"
)
//整数排序
func sortInte()  {
	var a =[...]int{12,23,4,2}//数字是值类型的所以不能直接传入
	sort.Ints(a[:])//a[:]代表是将a变成切片然后传入sort进行排序
	fmt.Print(a)
}

//数组排序
func sortString()  {
	var str =[...]string{"qwe","wefwv","fwcd"}
	sort.Strings(str[:])
	fmt.Println(str)
}

//浮点数排序
func sortFloat64()  {
	var fl64=[]float64{0.3,0.55,0.66}
	sort.Float64s(fl64[:])
	fmt.Println(fl64)
}

//搜索排序后的值
func searchInts()  {
	var intarr = [...]int{6,5,6,4,7,7,89,9,5,3,1,6}
	sort.Ints(intarr[:])//切记只有将数组的值类型转化成切片的值类型才能去排序不然直接Panic
	index :=sort.SearchInts(intarr[:],6)//搜索一定是完成排序之后才会搜索正确，当遇到重复的值的时候只会返回第一次出现的位置

	fmt.Println(index)
	fmt.Println(intarr)
}


func main() {
	sortInte()
	sortString()
	sortFloat64()
	searchInts()
}

