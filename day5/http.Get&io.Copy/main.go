package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	imgUrl := "https://www.twle.cn/static/i/img1.jpg"
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Printf("%#%s",err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("img1.jpg", data, 0644)
}