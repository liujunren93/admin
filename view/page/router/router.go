package router

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"os"
)

func BuildRouter(path string,group core.Group)  {
	path=fmt.Sprintf("%s/%s",path,global.FilePath[global.TypeHRouter])
	err := os.MkdirAll(path, 0766)
	if err != nil {
		panic(err)
	}
	var indexPage index
	for _, dom := range group.List {
		var routers routerList
		indexPage.add(fmt.Sprintf("import %s from '@%s\\%s.js'\n",dom.Name,global.FilePath[global.TypeRouter][3:],dom.Name),
			dom.Name)
		routers = append(routers,router{
			path:      "",
			name:      "",
			component: "",
			meta:      "",
			redirect:  "",
			children:  nil,
		} )

	}
}
