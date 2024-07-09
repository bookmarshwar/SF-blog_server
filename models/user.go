package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName  string `gorm:"size:20" json:"username"`
	NickName  string `gorm:"size:20" json:"nickname"`
	Telephone string `gorm:"szie:11" json:"telephone"`
	Password  string `gorm:"size:20" json:"password"`
	Email     string `gorm:"size:255" json:"email"`
}
