package model

import "gorm.io/gorm"


type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;default:""`
	Hash string `gorm:"type:char(56);not null;default:""`
}

type Url struct {
	gorm.Model
	Method string `gorm:"method;type:varchar(50);not null;default:''"`
	Path string `gorm:"path;type:varchar(50);not null;default:''"`

}