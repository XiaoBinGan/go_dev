package handler

import (
	"context"
	"github.com/prometheus/common/log"
	"go_dev/day18/e_commerce/category/common"
	"go_dev/day18/e_commerce/category/domain/model"
	"go_dev/day18/e_commerce/category/domain/service"
	category "go_dev/day18/e_commerce/category/proto/category"
)


type Category struct {
	CategoryDataService service.ICategoryDataService//绑定接口的类型
}
/**provide  create category server
   1.variable  &model.Category
   2.json marshal then unmarshal for category json tag
   3.use category.CategoryDataService ->CategoryRepository->gorm,db
   4.response the struct
   5.return err
 */
func (c *Category)CreateCategory(ctx context.Context, req *category.CategoryRequest, rep *category.CreateCategoryResponse) error {
	category :=&model.Category{}
	err := common.SwapTo(req, category)
	if err!=nil{
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err!=nil{
		return err
	}
	rep.Message="分类添加成功"
	rep.CategoryId=categoryId
	return nil
}
//provide update category server
func (c *Category)UpdateCategory(ctx context.Context,req *category.CategoryRequest,rep *category.UpdateCategoryResponse)error{
	category:= &model.Category{}
	err := common.SwapTo(req, category)
	if err!=nil{
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err!=nil{
		rep.Message="分类更新失败"
		return err
	}
	rep.Message="分类更新成功"
	return nil
}
//provide delete category server
func (c *Category) DeleteCategory(ctx context.Context,req *category.DeleteCategoryRequest, rep *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(req.CategoryId)
	if err!=nil{
		return err
	}
	rep.Message="删除成功"
	return nil
}
//provide Find category by category_name server
func(c *Category) FindCategoryByName(ctx context.Context,req *category.FindByNameRequest, rep *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err!=nil{
		return err
	}
	return common.SwapTo(category,rep)
}
//provide Find Category By ID server
func (c *Category)FindCategoryByID(ctx context.Context,req *category.FindByIDRequest,rep *category.CategoryResponse)error {
	categoryByID, err := c.CategoryDataService.FindCategoryByID(req.CategoryId)
	if err!=nil{
		return err
	}
	return common.SwapTo(categoryByID,rep)
}
//provide Find Category By Level server
func (c *Category)FindCategoryByLevel(ctx context.Context,req *category.FindByLevelRequest,rep *category.FindAllResponse)error{
	categoryByLevel, err := c.CategoryDataService.FindCategoryByLevel(req.Level)
	if err!=nil{
		return err
	}
	categoryToResponse(categoryByLevel,rep)
	return nil
}
//provide Find Category By Parent server
func (c *Category)FindCategoryByParent(ctx context.Context,req *category.FindByParentRequest,rep *category.FindAllResponse)error  {
	categoryByParent, err := c.CategoryDataService.FindCategoryByParent(req.ParentId)
	if err !=nil{
		return err
	}
	categoryToResponse(categoryByParent,rep)
	return nil
}
//provide Find All category server
func (c *Category)FindAllCategory(ctx context.Context,req *category.FindAllRequest,rep *category.FindAllResponse)error{
	allCategory, err := c.CategoryDataService.FindAllCategory()
	if err !=nil{
		return err
	}
	categoryToResponse(allCategory,rep)
	return nil
}

/**
  category to Response
  1.range category slice
  2.new categoryResponse type
  3.reflect json for model json tag
  4.append (FindAllResponse,category_item)
 */

func categoryToResponse(categorySlice []model.Category,response *category.FindAllResponse)  {
	for _, cg := range categorySlice {
		cr :=&category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err!=nil{
			log.Error(err)
			break
		}
		response.Category=append(response.Category,cr)
	}
}