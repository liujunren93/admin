package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Account  string `gorm:"account;type:varchar(100);not null;default:''"`
	Name     string `gorm:"name;type:varchar(100);not null;default:''"`
	Password string `gorm:"password;type:char(60);not null;default:''"`
}

type Role struct {
	gorm.Model
	Name   string `gorm:"name;type:varchar(100);not null;default:''"`
	Status int8   `gorm:"status;type:tinyint(1);not null;default:'1'"`
}

type RolePermission struct {
	gorm.Model
	PermissionID int `gorm:"permission_id;type:int;not null;default:''"`
	RoleID       int `gorm:"role_id;type:int;not null;default:''"`
}
