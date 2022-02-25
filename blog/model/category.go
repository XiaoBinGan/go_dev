package model
/**
分类的结构体对应myqsl分类表
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '分类Id',
  `category_name` varchar(255) COLLATE utf8mb4_0900_bin NOT NULL COMMENT '分类名称',
  `category_no` int NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_bin;
*/
type Categroy struct {
	CategoryID int64 `db:"id"`
	CategroyName string `db:"category_name"`
	CategroyNo int `db:"category_no"`
}