package handler

import (
	"context"
	"go_dev/day18/e_commerce/product/common"
	"go_dev/day18/e_commerce/product/domain/model"
	"go_dev/day18/e_commerce/product/domain/service"
	. "go_dev/day18/e_commerce/product/proto/product"
)

type Product struct {
    ProductDateService service.IProductDataService
}
//add product
func(p *Product)AddProduct(ctx context.Context,req *ProductInfo,res *ResponseProduct)  error{
	product :=&model.Product{}
	err := common.SwapTo(req, product)
	if err!=nil{
		return err
	}
	addProduct, err := p.ProductDateService.AddProduct(product)
	if err!=nil{
		return err
	}
	res.ProductId=addProduct
	return nil
}
//Find product by ID
func(p *Product)FindProductById(ctx context.Context,req *RequestID,res *ProductInfo) error{
	productByID, err := p.ProductDateService.FindProductByID(req.ProductId)
	if err!=nil{
		return err
	}
	if err := common.SwapTo(productByID, res);err!=nil{
		return err
	}
	return nil
}
//update product
func (p *Product)UpdateProduct(ctx context.Context,req *ProductInfo,res *Response) error {
	product :=&model.Product{}
	if err :=common.SwapTo(req,product);err!=nil{
		return err
	}
	if err := p.ProductDateService.UpdateProduct(product);err!=nil{
		return err
	}
	res.Msg="update success"
	return nil
}
//delete product
func(p *Product)DeleteProductByID(ctx context.Context,req *RequestID,res *Response) error{
	err := p.ProductDateService.DeleteProduct(req.ProductId)
	if err!=nil{
		return err
	}
	res.Msg="delete success"
	return nil
}
//find all product
func (p *Product)FindAllProduct(ctx context.Context,req *RequestAll,res *AllProduct) error{
	allProduct, err := p.ProductDateService.FindAllProduct()
	if err!=nil{
		return err
	}
	for _, v := range allProduct {
		productInfo :=&ProductInfo{}
		 if err := common.SwapTo(v, productInfo);err!=nil{
			 return nil
		 }
		res.ProductInfo=append(res.ProductInfo,productInfo)
	}
	return nil
}