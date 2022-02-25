package model

type Category struct{
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	CategoryName string `gorm:"unique_index，not_null"json:"category_name"`//unique_index唯一索引 not_null不能为空
	CategoryLevel uint32 `json:"category_leve"`
	CategoryParent int64 `json:"category_parent"`
	CategoryImage string `json:"category_image"`
	CategoryDescription string `json:"category_description"`
}

