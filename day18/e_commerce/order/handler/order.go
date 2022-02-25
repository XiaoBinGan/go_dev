package handler

import (
	"context"
	"go_dev/day18/e_commerce/order/domain/model"
	"go_dev/day18/e_commerce/order/domain/service"
      . "go_dev/day18/e_commerce/order/proto/order"
	"github.com/XiaoBinGan/common"

)
type Order struct {
	OrderService service.IOrderService
}


func(o *Order)GetOrderByID(ctx context.Context,req *OrderID,res *OrderInfo) error{
	order, err := o.OrderService.FindOrderByID(req.OrderId)
	if err!=nil{
		return err
	}
	return common.SwapTo(order, res)
}
func(o *Order)GetAllOrder(ctx context.Context,req *AllOrderRequest,res *AllOrder) error{
	orders, err := o.OrderService.FindAll()
	if err!=nil{
		return err
	}
	for _,v := range orders {
		order :=&OrderInfo{}
		if err := common.SwapTo(v, order);err!=nil{
			return err
		}
		res.OrderInfo=append(res.OrderInfo,order)
	}
	return nil
}
func(o *Order)CreateOrder(ctx context.Context,req *OrderInfo,res *OrderID) error{
	order :=&model.Order{}
	if err := common.SwapTo(req, order);err!=nil{
		return err
	}
	if addOrder, err := o.OrderService.AddOrder(order);err!=nil{
		return err
	}else {
		res.OrderId=addOrder
	}
	return nil
}
func(o *Order)DeleteOrderByID(ctx context.Context,req *OrderID,res *Response) error{
	if err := o.OrderService.DeleteOrderByID(req.OrderId);err!=nil{
		return err
	}
	res.Msg="order delete success"
	return nil
}
func(o *Order)UpdateOrderPayStatus(ctx context.Context,req *PayStatus,res *Response) error{
	if err := o.OrderService.UpdatePayStatus(req.OrderId, req.PayStatus);err!=nil{
		return err
	}
	res.Msg="the pay status changed success"
	return nil
}
func(o *Order)UpdateOrderShipStatus(ctx context.Context,req *ShipStatus,res *Response) error{
	if err :=o.OrderService.UpdateShipStatus(req.OrderId,req.ShipStatus);err!=nil{
		return err
	}
	res.Msg="the ship status changed success"
	return nil
}
func(o *Order)UpdateOrder(ctx context.Context,req *OrderInfo,res *Response) error{
	order :=&model.Order{}
	if err := common.SwapTo(req, order);err!=nil{
		return err
	}
	if err := o.OrderService.UpdateOrder(order);err!=nil{
		return err
	}
	res.Msg="update the order changed success"
	return nil
}