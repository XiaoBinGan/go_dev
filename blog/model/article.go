package model

import "time"

/**文章结构体 对应mySQL的文章表
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint NOT NULL COMMENT '分类Id',
  `content` longtext COLLATE utf8mb4_0900_bin NOT NULL COMMENT '文章内容',
  `title` varchar(1024) COLLATE utf8mb4_0900_bin NOT NULL COMMENT '文章标题',
  `view_count` int NOT NULL COMMENT '阅读次数',
  `comment_count` int NOT NULL COMMENT '评论次数',
  `username` varchar(255) COLLATE utf8mb4_0900_bin NOT NULL COMMENT '作者',
  `status` int NOT NULL DEFAULT '1' COMMENT '状态,正常为1',
  `summary` varchar(255) COLLATE utf8mb4_0900_bin NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin;

*/


type ArticleInfo struct {
	Id int64 			`db:"id"`
	CategroyId int64 	`db:"category_id"`
						//`db:"content"`文章内容冗余首页暂时不用
	Title string 		`db:"title"`
	ViewCount uint32  	`db:"view_count"`
	CommmentCount uint32`db:"comment_count"`
	UserName string		`db:"username"`
	Status int			`db:"status"`
	Summary string 		`db:"summary"`
	CreateTime time.Time `db:"create_time"`

}
//文章内容详情
//为了提升效率
type ArticleDetail struct {
	ArticleInfo
	Content string     `db:"content"`
	Categroy
}
//文章引到上下分页
type ArticleRecord struct {
	ArticleInfo
	Categroy
}