package server

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/global"
	"os"
	"strings"
)

func BuildSql(groups ...*core.Group)  {
	create, err := os.Create(global.FilePath[global.TypeDb]+"/rbac.sql")
	if err != nil {
		panic(err)
	}
	defer create.Close()
	sql:="insert into path(`name`,`method`,`path`)values"
	var sqls []string
	for _, group := range groups {
		for _, dom := range group.List {
			sqls = append(sqls, fmt.Sprintf("(%q,%q,%q)",dom.HName+"列表","Get",dom.Name))
			sqls = append(sqls, fmt.Sprintf("(%q,%q,%q)",dom.HName+"删除","Delete",dom.Name))
			sqls = append(sqls, fmt.Sprintf("(%q,%q,%q)",dom.HName+"创建","Post",dom.Name))
			sqls = append(sqls, fmt.Sprintf("(%q,%q,%q)",dom.HName+"编辑","Put",dom.Name))
		}
	}
	sql+=strings.Join(sqls, ",")
	create.WriteString(sql)

}