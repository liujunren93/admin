package table

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/core"
	"os"
	"strings"
)
//go:embed table.vue
var table string


func NewPage(path string,groups []*core.Group) {
	os.MkdirAll(path,0766)
	for _, group := range groups {
		for _, dom := range group.List {
			p := parseDom(*dom)
			buildFile(path,p)
		}

	}


}

func buildFile(path string,p page)  {
	create, err := os.Create(path + "/" + p.name + ".vue")
	defer create.Close()
	if err != nil {
		panic(err)
	}

	sprintf := fmt.Sprintf(table, p.search, p.topOperator, p.table,strings.Join(p.importStr, "\n"),p.column,
		strings.Join(p.data, "\n"))
	create.WriteString(sprintf)

}