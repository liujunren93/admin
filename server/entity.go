package server

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"github.com/liujunren93/admin/utils"
)

func NewEntity(gs ...*core.Group) {
	var entity []File
	for _, g := range gs {
		var f File
		f.Name = g.Name
		f.Pkg = "entity"
		f.Import = append(f.Import, global.Mod+"/"+global.FilePath[global.TypeModel])
		for _, dom := range g.List {
			f.Class = append(f.Class,entityList(dom)... )
		}
		entity = append(entity, f)
	}
	Build(global.TypeEntity,entity...)
}

func entityList(dom *core.Dom) []Class {
	var reqClass, resClass Class
	reqClass.Name=dom.Name+"ListReq"
	resClass.Name=dom.Name+"ListRes"
	if dom.HPagination {
		reqClass.AddField("Page", "int", "from:\"page\"")
		resClass.AddField("Total", "int", "json:\"total\"")

	}
	resClass.AddField("List", fmt.Sprintf("[]model.%s", dom.Name), "json:\"list\"")
	if dom.HasSearch {
		for _, field := range dom.Fields {
			if field.HSearch == "" {
				continue
			}
			var tag = field.Tag
			tag = append(tag, fmt.Sprintf("from:\"%s\"", utils.SnakeString(field.Name)))
			reqClass.AddField(field.Name, field.Type, tag...)
		}

	}
	return []Class{reqClass, resClass}
}
