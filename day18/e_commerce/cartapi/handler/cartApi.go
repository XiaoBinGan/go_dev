package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cart "github.com/XiaoBinGan/cart/proto/cart"
	"github.com/micro/go-micro/v2/util/log"
	cartApi "go_dev/day18/e_commerce/cartapi/proto/cartApi"
	"net/http"
	"strconv"
)
type CartApi struct {
	CartService cart.CartService
}



//CartApi.Call 通过API向外暴露/cartApi/call .接收http请求
//即 /cartAPI/call请求会调用go.micro.api.cartApi 服务的Cart.Call方法
func (c *CartApi)FindAll(ctx context.Context,req *cartApi.Request,res *cartApi.Response) error{
   log.Info("接收到/cartApi/findAll 访问请求")
   //Whether a parameter of user id is received
	if _,ok := req.Get["user_id"];!ok{
		res.StatusCode=500
		return errors.New("参数异常")
	}
	//get user_id values
	userIdString := req.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	//Determine whether the parameters are correct
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err!=nil{
		return err
	}
	//to obtain all of the product on the cart
	cartAll, err := c.CartService.GetAll(context.TODO(), &cart.CartFindAll{
		UserId: userId,
	})
	//marshal the parameters  fo web
	bytes, err := json.Marshal(cartAll)
	if err!=nil{
		return err
	}
	//have not err ,return success status and response successful result
	res.StatusCode=http.StatusOK
	res.Body=string(bytes)
	return err
}


//Ctrl+A：到行首（达到Home键的效果） Ctrl+E：到行尾（达到End键的效果） Ctrl+N：到下一行Ctrl+P：到上一行