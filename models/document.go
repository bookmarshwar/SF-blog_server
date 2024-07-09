package models

import (
	"gorm.io/gorm"
)

type DocumentModel struct {
	gorm.Model
	Briefly string `gorm:"size:255" json:"briefly"`      //摘要
	Title   string `gorm:"size:256" json:"title"`        //标题
	Cover   string `gorm:"size:256" json:"cover"`        //封面
	Path    string `gorm:"size:256" json:"documentPath"` //路径
	//AuthorId   int       `gorm:""`          //作者id
	//Tag
}
