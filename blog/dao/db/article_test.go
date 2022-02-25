package db

import (
	"go_dev/blog/model"
	"testing"
	"time"
)

func init() {
	dns :="root:@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init("mysql", dns)
	if err!=nil{
		panic(err)
	}
}



func TestInsertArticle(t *testing.T) {
	article :=&model.ArticleDetail{
	}
	article.ArticleInfo.CategroyId=1
	article.ArticleInfo.Title="html how to learn"
	article.ArticleInfo.ViewCount=200
	article.ArticleInfo.CommmentCount=1
	article.ArticleInfo.UserName="吴佳浩"
	article.ArticleInfo.Summary="值得学习的文章"
	article.ArticleInfo.Status=1
	article.ArticleInfo.CreateTime = time.Now()
	article.Content="文章的内容非常多这里省略1万字"
	article.Categroy.CategroyName="html"
	article.Categroy.CategroyNo=1
	articleId, err := InsertArticle(article)
	if err!=nil{
		return
	}
	t.Logf("articleId:%+#v",articleId)

}