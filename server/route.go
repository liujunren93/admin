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
	for _, group := range groups {
		for _, dom := range group.List {
			routeName:=utils.UcFirst(dom.Name)
			funcBody.WriteString(fmt.Sprintf("\t%s:=engine.Group(\"%s\")\n",routeName,routeName))
			funcBody.WriteString("\t{\n")
			funcBody.WriteString(fmt.Sprintf("\t\t%s.Get(\"\",%s.%sList)\n",routeName,group.Name+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.Post(\"\",%s.%sCreate)\n",routeName,group.Name+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.Put(\"/:id\",%s.%sUpdate)\n",routeName,group.Name+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.Put(\"/:id\",%s.%sInfo)\n",routeName,group.Name+"Ctrl()",dom.Name))
			funcBody.WriteString(fmt.Sprintf("\t\t%s.Delete(\"/:id\",%s.%sDel)\n",routeName,group.Name+"Ctrl()",dom.Name))
			funcBody.WriteString("\t}\n")
		}
	}
	funcBody.WriteString("return engine")
	route.AddFunc(NewFunc("Router","","engine *gin.Engin",funcBody.String()))
	Build(global.TypeRouter,route)
}
