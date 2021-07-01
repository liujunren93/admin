package static

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

type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null;default:""`
	Hash string `gorm:"type:char(56);not null;default:""`
}
type PermissionUrl struct {
	gorm.Model
	PermissionID int `gorm:"permission_id;type:int;not null;default:0"`
	UrlID        int `gorm:"url_id;type:int;not null;default:0"`
}

type Url struct {
	gorm.Model
	Name   string `gorm:"name;type:varchar(100);not null;default:''"`
	Method string `gorm:"method;type:varchar(50);not null;default:''"`
	Path   string `gorm:"path;type:varchar(50);not null;default:''"`
}

type Menu struct {
	gorm.Model
	ParentID  int    `gorm:"parent_id;type:int;not null;default:0"`
	Title     string `gorm:"title;type:varchar(100);not null;default:''"`
	Key       string `gorm:"key;type:char(56);not null;default:''"`
	Component string `gorm:"component;type:varchar(100);not null;default:''"`
	Icon      string `gorm:"icon;type:varchar(100);not null;default:''"`
	Meta      string `gorm:"meta;type:varchar(255);not null;default:''"`
}
