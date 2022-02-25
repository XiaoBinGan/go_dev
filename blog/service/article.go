package service

import (
	"go_dev/blog/dao/db"
	"go_dev/blog/model"
)
//get All articleRecord
func GetArticleRecordList(pageNum,pageSize int)(articleRecordList []*model.ArticleRecord,err error) {
	//1.get articleList
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err!=nil ||len(articleInfoList)<=0{
		return
	} 

	//2.get article's all category
	ids := getCategoryIds(articleInfoList)
	Categorylist, err := db.GetCategoryList(ids)
	if err!=nil{
		return
	}
//range all list for gw
	for _, article:= range articleInfoList{
		articleRecord :=&model.ArticleRecord{
			ArticleInfo:*article,
		}
		categoryId :=article.CategroyId
		for _, category := range Categorylist {
			if categoryId== category.CategoryID{
				articleRecord.Categroy=*category
				break
			}
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}

//get the category article and match category information, by the categroyId
func GetArticleRecordListById(catagoryId,pageNum,pageSize int)(articleRecordList []*model.ArticleRecord,err error){
	//1.get AllarticleInfoList by category and pageNum and pageSize
	articleInfoList, err := db.GetArticleListByCategoryId(catagoryId, pageNum, pageSize)
	if err!=nil ||len(articleInfoList)<=0{
		return
	}
	//2.get All category by articleInfoList's idList
	ids :=getCategoryIds(articleInfoList)
	//3.get the articleInfoList's categoryIdList
	categoryList, err := db.GetCategoryList(ids)
	if err!=nil{
		return
	}
	//4.convergence parameters for gateway
	for _,article := range articleInfoList {
		articleRecord :=&model.ArticleRecord{
			ArticleInfo:*article,
		}
		categoryId :=article.CategroyId
		for _, category := range categoryList {
			if categoryId== category.CategoryID{
				articleRecord.Categroy=*category
				break
			}
		}
		articleRecordList = append(articleRecordList,articleRecord)
	}
	return
}

//Id based on multiple articles, Gets a collection of multiple classification
func getCategoryIds(articleInfoList []*model.ArticleInfo)(ids []int64){
	for _, article := range articleInfoList {
		categoryId :=article.CategroyId
		for _, id := range ids {
			if id!=categoryId{
				ids = append(ids,categoryId)
			}
		}
	}
	return
}