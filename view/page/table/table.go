package table

import (
	"fmt"
	"github.com/liujunren93/admin/core"
)

var table string

func NewPage(groups []*core.Group) {
	var pageList []page
	for _, group := range groups {
		for _, dom := range group.List {
			pageList = append(pageList, parseDom(*dom))
		}

	}
fmt.Println(pageList)

}
