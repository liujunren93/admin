package test

import "gorm.io/gorm"
//@curd:name=管理员
//@curd:page=管理员
type Banner struct {
	gorm.Model
	//@curd:name=管理员;search=like;sort=1
	Name string `gorm:"type:varchar(100)" json:"name"`
	Password string  `gorm:"type:varchar(100)"json:"password"`
}

func (Banner)List(string2 string)(str string,err error)  {
return "", err
}
func (Banner)Info(string2 string)(str string,err error)  {
	return "", err
}

//@curd:name=详情
type BannerInfo struct {
	NickName string `gorm:"type:varchar(100)" json:"nick_name"`
}
