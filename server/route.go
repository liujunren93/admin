package server

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"github.com/liujunren93/admin/utils"
	"strings"
)

func NewRoute(groups ...*core.Group) {
	var route File
	route.Pkg="router"
	route.Name="router"
	route.Import = append(route.Import, "github.com/gin-gonic/gin", global.Mod+"/"+global.FilePath[global.TypeCtrl])
	var funcBody strings.Builder
	funcBody.WriteString("engine=gin.Default()\n")
	for _, group := range groups {

		for _, dom := range group.List {
			groupName:=utils.UpFirst(group.Name)
			routeName:=utils.UcFirst(dom.Name)

			funcBody.WriteString(fmt.Sprintf("\t%s:=engine.Group(\"%s\")\n",routeName,routeName))
			funcBody.WriteString("\t{\n")

			funcBody.WriteString(fmt.Sprintf("\t\t%s.GET(\"\",ctrl.%s.%sList)\n",routeName,groupName+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.POST(\"\",ctrl.%s.%sCreate)\n",routeName,groupName+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.PUT(\"/:id\",ctrl.%s.%sUpdate)\n",routeName,groupName+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.PUT(\"/:id\",ctrl.%s.%sInfo)\n",routeName,groupName+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.DELETE(\"/:id\",ctrl.%s.%sDelete)\n",routeName,groupName+"Ctrl()",dom.Name))
			funcBody.WriteString("\t}\n")
		}
	}
	funcBody.WriteString("return engine")
	route.AddFunc(NewFunc("Router","","engine *gin.Engine",funcBody.String()))
	Build(global.TypeRouter,route)
}
