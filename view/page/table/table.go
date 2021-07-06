package table

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"os"
	"strings"
)
//go:embed table.vue
var table string


func BuildPage(groups []*core.Group) {
	path:=global.WebRoot
	for _, group := range groups {
		for _, dom := range group.List {
			p := parseDom(*dom)
			buildFile(path,dom.Name,p)
		}

	}


}

func buildFile(path,name string,p page)  {
	tablePath:=fmt.Sprintf("%s/%s/%s/",path,global.FilePath[global.TypeHView],name)
	os.MkdirAll(tablePath,0766)
	create, err := os.Create(tablePath+"/List.vue")
	defer create.Close()
	if err != nil {
		panic(err)
	}

	sprintf := fmt.Sprintf(table, p.search, p.topOperator, p.table,strings.Join(p.importStr, "\n"),p.column,
		strings.Join(p.data, "\n"))
	create.WriteString(sprintf)

}