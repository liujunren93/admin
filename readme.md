## 一个根据struct文件生成后端api 和前端页面的工具

---

### struct
+ 注释介绍：
  + struct
      1. //@curd :本工具会使用的注释
      2. name=前端页面名字;page=是否分页（Boolean）
  + struct.field
    1. name: 页面显示名字
    2. search 是否支持搜索
    3. type： 页面显示类型
       1. select select([{"label":"label","value":"value"},{"label":"label1","value":"value1"}])
       2. checkbox
       3. editor
       4. input
       5. radio
       6. textarea
       7. upload
    4.sort:是否用于排序
+ 列子：
```go
//@curd:name=管理员;page=管理员 
type Admin struct {
	gorm.Model
	//@curd:name=管理员;search=true;type=select(banner);sort=1
	Name string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	//@curd:name=密码;search=true;type=select([{"label":"label","value":"value"},{"label":"label1","value":"value1"}]);sort=2
	Password string  `gorm:"type:varchar(100)"json:"password"`
}

```