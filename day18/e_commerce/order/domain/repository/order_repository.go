package repository

import (
	"errors"
	"github.com/micro/go-micro/v2/util/log"
	"go_dev/day18/e_commerce/order/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type IOrderRepository interface {
	InitTable()error
	FindOrderByID(int64)(*model.Order,error)
	FindAll()([]model.Order,error)
	CreateOrder(*model.Order)(int64,error)
	DeleteOrderByID(int64)error
    UpdateOrder(*model.Order)error
	UpdatePayStatus(int64,int32)error  //order_id changeNumber
	UpdateShipStatus(int64,int32)error
}

//new order Repository
func NewOrderRepository(db *gorm.DB) IOrderRepository  {
	return &OrderRepository{mysqlDb:db}
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}
func (o *OrderRepository)InitTable()error{
	return o.mysqlDb.CreateTable(&model.Order{},&model.OrderDetail{}).Error
}
func (o *OrderRepository)FindOrderByID(orderId int64)(order *model.Order,err error){
	order =&model.Order{}
	return order,o.mysqlDb.Preload("OrderDetail").First(order,orderId).Error
}
func (o *OrderRepository)FindAll()(orderAll []model.Order,err error){
	//Preload("OrderDetail").Find(&orderAll)//find  orderDetail by out key for result
	return orderAll,o.mysqlDb.Preload("OrederDetail").Find(orderAll).Error
}
func(o *OrderRepository)CreateOrder(order *model.Order)(orderID int64,err error){
       return order.ID,o.mysqlDb.Create(order).Error
	}
func(o *OrderRepository)DeleteOrderByID(orderId int64)error{
	t := o.mysqlDb.Begin()
	//well  return need recover err
	defer func() {
		if r := recover();r!=nil{
			t.Rollback()
			log.Error(r)
		}
	}()
	//catch begin error
	if t.Error!=nil{
		return nil
	}
	//delete order  info
	if err := o.mysqlDb.Unscoped().Where("order_Id=?", orderId).Delete(&model.Order{}).Error;err!=nil{
		return err
	}
	//delete order detail
	if err:=o.mysqlDb.Unscoped().Where("order_id = ?",orderId).Delete(&model.OrderDetail{}).Error;err!=nil{
		return err
	}
	//begin end commit & catch error
	return  t.Commit().Error
}
func(o *OrderRepository)UpdateOrder(order *model.Order)error{
	return o.mysqlDb.Model(order).Update(order).Error
}
func(o *OrderRepository)UpdatePayStatus(ID int64,changeNumber int32)error{
	db := o.mysqlDb.Model(&model.Order{}).Where("id = ?", ID).UpdateColumn("pay_status", changeNumber)
	if db.Error!=nil{
		return db.Error
	}
	if db.RowsAffected==0{
		return errors.New("pay status update failed")
	}
	return nil
}
func(o *OrderRepository)UpdateShipStatus(ID int64,shipCode int32)error{
	db := o.mysqlDb.Model(&model.Order{}).Where("id =?", ID).UpdateColumn("ship_status", shipCode)
	if db.Error!=nil{
		return db.Error
	}
	if db.RowsAffected==0{
		return errors.New("update ship status is failed,please try can later")
	}
   return nil
}
