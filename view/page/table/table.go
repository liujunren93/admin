package table

import (
	"fmt"
	"github.com/liujunren93/admin/core"
)

var table string

func NewPage(dom core.Dom) {
	p := parseDom(dom)

	fmt.Printf("%#v",p)

}
