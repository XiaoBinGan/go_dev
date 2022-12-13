package db

import (
	"blog/model"

	_ "github.com/go-sql-driver/mysql"
)

/*
*
InsertArticle insert single article
*/
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//if article == nil {
	//	return
	//}
	sqlstr := `insert into 
				article(content,summary,title,username,category_id,view_count,comment_count)
				values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, article.Content,
		article.Summary,
		article.Title,
		article.UserName,
		article.CategroyId,
		article.ViewCount,
		article.CommmentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

/*
*
GetArticleList get article list
*/
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlstr := `select id,category_id,titleview_count,comment_count,username,status,summary,
			 from article  
			 where status = 1
			 order by create_time desc 
             limit ?,?`
	err = DB.Select(&articleList, sqlstr, pageNum, pageSize)
	return
}

/*
*
GetArticleDetail
*/
func GetArticleDetail(articleID int64) (articleDetail model.ArticleDetail, err error) {
	sqlstr := `select id,category_id,titleview_count,comment_count,username,status,summary,content
		      from article
			  where id =? and status = 1`
	err = DB.Select(&articleDetail, sqlstr, articleID)
	return
}

func GetArticleListByCategoryId(category, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	sqlstr := `select id,category_id,titleview_count,comment_count,username,status,summary
				from article
				where status=1
				and category_id=?
				order by create_time desc
				limit ?,?`
	err = DB.Select(&articleList, sqlstr, category, pageNum, pageSize)
	return
}
