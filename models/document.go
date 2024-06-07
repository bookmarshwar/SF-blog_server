package models

import "time"

type DocumentModel struct {
	DocID      int       //id
	Briefly    string    //摘要
	Title      string    //标题
	Cover      string    //封面
	CreateTime time.Time //创建时间
	Path       string    //路径
	AuthorId   int       //作者id
}
