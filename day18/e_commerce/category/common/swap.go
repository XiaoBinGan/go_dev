package common

import "encoding/json"

func SwapTo(request,category interface{})(err error){
	dataByte, err := json.Marshal(request) //json编码请求的参数
	if err!=nil{
		return err
	}
	return json.Unmarshal(dataByte,category) //解码json对应到数据库的json tag
}