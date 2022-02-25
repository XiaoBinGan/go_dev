package ServiceImpl

import (
	"context"
	"go_dev/day18/go-micro-grpc/Services"
	"strconv"
)



//Test function
func newProd(id int32,pname string)*Services.ProdModel  {
	return &Services.ProdModel{ProdId: id,ProdName: pname}
}

type ProdService struct {}


//不需要对修改过的rep进行手动返回 直接修改rep即刻
func (p *ProdService)GetProdsList(ctx context.Context, req *Services.ProdsRequest, rep *Services.ProdListResponse) error  {
	models :=make([]*Services.ProdModel,0)
	var i int32
	for i=0;i<req.Size;i++ {
		models=append(models,newProd(100+i,"prodname"+strconv.Itoa(100+int(i))))
	}
	rep.Data=models
	return nil
}