package test

import "gorm.io/gorm"
//@curd:name=管理员;page=管理员
type Admin struct {
	gorm.Model
	//@curd:name=管理员;search=true;type=select(banner);sort=1
	Name string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	//@curd:name=密码;search=true;type=select([{"label":"label","value":"value"},{"label":"label1","value":"value1"}]);sort=2
	Password string  `gorm:"type:varchar(100)"json:"password"`
}

func (Admin)Login(string2 string)(str string,err error)  {
return "", err
}
func (Admin)Login1(string2 string)(str string,err error)  {
	return "", err
}
func Login(string2 string)(str string,err error)  {
	return "", err
}
//@curd:name=详情
type AdminInfo struct {
	NickName string `gorm:"type:varchar(100)" json:"nick_name"`
}
