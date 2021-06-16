package Info

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"os"
	"strings"
)

//go:embed info.vue
var info string
func BuildInfo(path string,groups []*core.Group)  {

	for _, group := range groups {
		for _, dom := range group.List {
			filePath:=path+"/"+global.FilePath[global.TypeHView]+"/" + dom.Name
			err := os.MkdirAll(filePath, 0766)
			if err != nil {
				panic(err)
			}
			create, err := os.Create(filePath+ "/Info.vue")
			if err != nil {
				panic(err)
			}
			p := parseDom(*dom)
			sprintf := fmt.Sprintf(info, p.formItem, strings.Join(p.importList, ",\n"), strings.Join(p.components, ",\n"), strings.Join(p.data, ",\n"))
			_, err = create.WriteString(sprintf)
			if err != nil {
				panic(err)
			}
			create.Close()
		}
	}

}