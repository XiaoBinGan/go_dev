package service

import (
	"go_dev/day18/e_commerce/order/domain/model"
	"go_dev/day18/e_commerce/order/domain/repository"
)

type IOrderService interface {
	FindOrderByID(int64)(*model.Order,error)
	FindAll()([]model.Order,error)
	AddOrder(*model.Order)(int64,error)
	DeleteOrderByID(int64)error
	UpdateOrder(*model.Order)error
	UpdatePayStatus(int64,int32)error  //order_id changeNumber
	UpdateShipStatus(int64,int32)error
}

type OrderService struct {
	OrderRepository repository.IOrderRepository
}

//create OrderService
func NewOrderService(OrderRepository repository.IOrderRepository)IOrderService{
	return &OrderService{OrderRepository:OrderRepository}
}
func(o *OrderService)AddOrder(order *model.Order)(orderID int64,err error){
	return o.OrderRepository.CreateOrder(order)
}
func(o *OrderService)DeleteOrderByID(orderID int64)error{
   return o.OrderRepository.DeleteOrderByID(orderID)
}
func(o *OrderService)UpdateOrder(order *model.Order)error{
	return o.OrderRepository.UpdateOrder(order)
}
func(o *OrderService)UpdatePayStatus(id int64,payCode int32)error { //order_id changeNumber
    return o.OrderRepository.UpdatePayStatus(id,payCode)
}
func(o *OrderService)UpdateShipStatus(id int64,shipCode int32)error{
    return o.OrderRepository.UpdateShipStatus(id,shipCode)
}
func(o *OrderService)FindOrderByID(orderId int64)(*model.Order,error){
  	return o.OrderRepository.FindOrderByID(orderId)
}
func(o *OrderService)FindAll()([]model.Order,error){
   return o.OrderRepository.FindAll()
}
