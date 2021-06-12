package server

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"github.com/liujunren93/admin/utils"
	"strings"
)

func NewDao(gs ...*core.Group) {
	var dao []File
	for _, g := range gs {
		for _, dom := range g.List {
			var f File
			className := utils.UcFirst(dom.Name)
			f.Pkg = "dao"
			f.Name = className
			class := Class{
				Name: className + "Dao",
			}
			f.AddVariable(fmt.Sprintf("var %s *%s \n",utils.UcFirst(dom.Name), className+"Dao"))
			f.AddVariable(fmt.Sprintf("var %sOne sync.Once \n", utils.UcFirst(dom.Name)))
			f.AddFunc(NewFunc(utils.UpFirst(class.Name), "", "*"+utils.UcFirst(class.Name), daoConstructBody(className, className+"Dao")))
			var searchBuf strings.Builder
			if dom.HPagination || dom.HasSearch {
				searchBuf = buildSearchBuf(dom)
				class.AddFunc(NewFunc("List", fmt.Sprintf("req entity.%sListReq", dom.Name), fmt.Sprintf("res entity.%sListRes", dom.Name), listDaoStr(searchBuf, dom.Name, searchBuf.Len() > 0)))
			} else {
				class.AddFunc(NewFunc("List", "", fmt.Sprintf("req entity.%sListRes", dom.Name), listDaoStr(searchBuf, dom.Name, false)))
			}
			class.AddFunc(NewFunc("Info", "id uint", fmt.Sprintf("data model.%s,err error", dom.Name), infoDaoStr()))
			class.AddFunc(NewFunc("Create", fmt.Sprintf("req model.%s", dom.Name), fmt.Sprintf("err error"), createDaoStr()))
			class.AddFunc(NewFunc("Update", fmt.Sprintf("id uint,req model.%s", dom.Name), fmt.Sprintf("err error"), upDaoStr()))
			class.AddFunc(NewFunc("Delete", "id uint", fmt.Sprintf("err error"), delDaoStr(dom.Name)))
			f.Import = append(f.Import,
				"sync",
				global.Mod+"/"+ global.FilePath[global.TypeDb],
				global.Mod+"/"+ global.FilePath[global.TypeModel],
				global.Mod+"/"+global.FilePath[global.TypeEntity] )
			f.Class = append(f.Class, class)
			dao = append(dao, f)

		}
	}
	Build(global.TypeDao,dao...)
}
func buildSearchBuf(dom *core.Dom) strings.Builder {
	var buf strings.Builder
	for _, field := range dom.Fields {
		if field.HSearch == "" {
			continue
		}
		if field.HSearch == "like" {
			buf.WriteString(fmt.Sprintf("\tif req.%s!=\"\" {\n", field.Name))
			buf.WriteString(fmt.Sprintf("\t\tDb=Db.where(\"%s like ?\",\"%%\"+req.%s+\"%%\")", utils.SnakeString(field.Name), field.Name))
		} else {
			buf.WriteString(fmt.Sprintf("\t\tDb=Db.Where(\"%s = ?\",req.%s)\n", utils.SnakeString(field.Name), field.Name))
		}
		buf.WriteString("\n\t}\n")
	}

	return buf
}
func daoConstructBody(variable, className string) string {
	code := `%sOne.Do(func(){
		%s=new(%s)
	})
	return %s
`
	return fmt.Sprintf(code, variable, variable, className, variable)
}

func listDaoStr(search strings.Builder, className string, pagination bool) string {

	var bodyStr strings.Builder
	if !pagination {
		bodyStr.WriteString(fmt.Sprintf("var list []model.%s \n", className))

	}
	bodyStr.WriteString("Db:=db.Db\n")

	bodyStr.WriteString(search.String())
	if pagination {
		bodyStr.WriteString(fmt.Sprintf("\tDb.Limit(offset).Offset(offset*req.Page).Find(&list) \n"))
		bodyStr.WriteString("\tDb.Count(&res.Total) \n")
		bodyStr.WriteString("\tres.List=list \n")
		bodyStr.WriteString("\treturn  res\n")
	} else {
		bodyStr.WriteString("\tDb.Find(&list) \n")
		bodyStr.WriteString("\treturn  list\n")
	}

	return bodyStr.String()
}

func infoDaoStr() string {
	bodyStr := `err=db.Db.Where("id=?",id).First(&data).Error
	return `
	return bodyStr
}
func createDaoStr() string {
	bodyStr := `err=db.Db.Create(&data).Error 
	return `
	return bodyStr
}

func upDaoStr() string {
	bodyStr := `err=db.Db.Where("id=?",id).Updates(&data).Error
	return `
	return bodyStr
}

func delDaoStr(className string) string {
	bodyStr := `err=db.Db.Where("id=?",id).Delete(model.%s{}).Error
	return `
	return fmt.Sprintf(bodyStr, className)
}
