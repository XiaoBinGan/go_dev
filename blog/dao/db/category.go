package db

import (
	"github.com/jmoiron/sqlx"
	"go_dev/blog/model"
)

/**
InitCategory insert into category(categoru_name,category_no)value(?,?)
*/
func InitCategory(c *model.Categroy)(CategoryId int64 ,err error) {
	sqlstr :="insert into category(category_name,category_no)value (?,?)"
	result, err := DB.Exec(sqlstr, c.CategroyName, c.CategroyNo)
	if err!=nil{
		return
	}
	CategoryId, err= result.LastInsertId()
	return
}

/**
	get single category history
	select id,category_name,category_no from category where id =?
	err=DB.Get(got value struct,Sql ,proviso)
 */
func GetCategoryById(id int64)(category *model.Categroy,err error) {
	category = &model.Categroy{}
	sqlstr:="select id,category_name,category_no from category where id =(?)"
	err = DB.Get(category, sqlstr, id)
	return
}

/**
GetCategoryList query rows of  the category

*/
func GetCategoryList(categoryIds []int64)(catrgoryList []*model.Categroy,err error)  {
	sqlstr, args, err := sqlx.In("select id,category_name,category_no from category where id in (?)", categoryIds)
	if err!=nil{
		return
	}
	err = DB.Select(&catrgoryList, sqlstr, args...)
	return
}

/**
	get all category
*/
func GetAllcategoryList()(categoryList []*model.Categroy,err error){
	sqlstr :="select id,category_name,category_no from category order by category_no asc"
	err  = DB.Select(&categoryList,sqlstr)
	return
}


















