package main

import (
	"encoding/json"
	"fmt"
	dat "go_dev/dataGw/data"
	"go_dev/dataGw/sorcket"
)
var pbsend interface{}

func main()  {
	//接口一
	var f *dat.Tasting
	err := json.Unmarshal([]byte(dat.A),&f)
	if err!=nil{
		fmt.Println("json unmarshal failed!")
		return
	}
	pbsend, err := json.Marshal(f)

	fmt.Print(pbsend)
	//for _,v := range f.Data {
		//fmt.Printf("%#+v\n",v)
	//}





	//接口二
	xd :=make(dat.XDataMap,20)
	da :=make(dat.DataMap,60)
	lp :=&dat.List{}
	//ap :=&dat.AllMap{}
	da["data"]=make(map[string][]int,40)
	xd["xData"]=[]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	da["data"]["邮件营销"]=[]int{120, 132, 101, 134, 90, 230, 210}
	da["data"]["联盟广告"]=[]int{220, 182, 191, 234, 290, 330, 310}
	da["data"]["视频广告"]=[]int{150, 232, 201, 154, 190, 330, 410}
	da["data"]["直接访问"]=[]int{320, 332, 301, 334, 390, 330, 320}
	da["data"]["搜索引擎"]=[]int{820, 932, 901, 934, 1290, 1330, 132}
	lp.XDataMap=xd
	lp.DataMap=da
	//ap.All=append(ap.All, lp)
	//fmt.Printf("%#v",ap.All[0])
	//for i, v := range xd {
	//	fmt.Println(i)
	//	for k, n := range v {
	//		fmt.Println("\t",k,n)
	//	}
	//}\
	//marshal, err := json.Marshal(ap)
	//fmt.Print(string(marshal))
	//err = json.Unmarshal([]byte(dat.B), ap)
	//if err!=nil{
	//	return
	//}
	ws := sorcket.NewWsServer(pbsend)
	ws.Start()
	//r :=gin.Default()
	//r.POST("/", func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK,gin.H{
	//		"status":1,
	//		"massage":"ok",
	//		"data":lp,
	//	})
	//	return
	//})
	//r.Run(":8080")
}