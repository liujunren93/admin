package table

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/utils"
	"github.com/liujunren93/admin/view/base"
	"github.com/liujunren93/admin/view/component"
	"github.com/liujunren93/admin/view/component/input"
	_select "github.com/liujunren93/admin/view/component/select"
	"regexp"
	"strings"
)

type page struct {
	search      string
	topOperator string
	table       string
	column      string
	data        []string
	importStr   []string
	method      []string
}
const searchItem = "<a-form-item label=\"%s\"> \n %s \n \t\t\t</a-form-item>\n"
func buildSearchForm(componenters ...component.Componenter) string {

	if len(componenters) == 0 {
		return ""
	}
	var t base.T = "\t"
	buf := strings.Builder{}
	bufAdvanced := strings.Builder{}
	for i, componenter := range componenters {
		if i < 2 {
			buf.WriteString(fmt.Sprintf("%s<a-col :md=\"8\" :sm=\"24\">", t.Multiple(2)))
			buf.WriteString("\n" + t.Multiple(3))
			buf.WriteString(fmt.Sprintf(searchItem, componenter.GetLabel(), componenter.Html(t.Multiple(4))))
			buf.WriteString(t.Multiple(2) + "</a-col>\n")
		} else {
			if bufAdvanced.Len() == 0 {
				bufAdvanced.WriteString(t.Multiple(2) + "<template v-if=\"advanced\">\n")
			}
			bufAdvanced.WriteString(t.Multiple(3) + "<a-col :md=\"8\" :sm=\"24\">")
			bufAdvanced.WriteString(fmt.Sprintf(searchItem, componenter.GetLabel(), componenter.Html(t.Multiple(4))))
			bufAdvanced.WriteString(t.Multiple(3) + "</a-col>\n")
		}
	}
	if bufAdvanced.Len() > 0 {
		bufAdvanced.WriteString(t.Multiple(2) + "</template>\n")
	}
	buf.WriteString(bufAdvanced.String())
	return buf.String()
}

//头部操作
func buildTopOperator(permission string) string {
	return fmt.Sprintf("\t"+`<a-button type="primary" icon="plus"  v-if="$auth('%s.add')" @click="handleAdd">新建</a-button>
        <a-dropdown v-action:edit v-if="selectedRowKeys.length > 0">
          <a-menu slot="overlay">
            <a-menu-item key="1"  v-if="$auth('%s.delBatch')"><a-icon type="delete" />删除</a-menu-item>
          </a-menu>
          <a-button style="margin-left: 8px">
            批量操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>`, permission, permission)
}
func getSearchComponent(page *page, dom core.Dom) {
	var components []component.Componenter
	page.importStr = append(page.importStr, fmt.Sprintf("import { getList } from '@/api/%s'", dom.Name) )
	for _, field := range dom.Fields {

		compile := regexp.MustCompile("select\\((\\w*)\\)")
		findString := compile.FindStringSubmatch(field.HType)

		if len(findString) > 1 {
			selectName := findString[1]
			components = append(components, _select.NewSimple(field.Name, field.HName, "queryParam."+field.Name, "", false, selectName+"Data"))
			if !json.Valid([]byte(selectName)) {
				upSelectName:=utils.UpFirst(selectName)
				page.importStr = append(page.importStr, fmt.Sprintf("import { getList as get%sList } from '@/api/%s'", upSelectName, upSelectName))
				page.data = append(page.data, fmt.Sprintf(`   %sData: () => {
        return  get%sList ()
          .then(res => {
            return res.data
          })
      },`, selectName, upSelectName))
			}
		}else if field.HSearch != "" {
			components = append(components, input.NewInput(field.Name, field.HName, "queryParam."+field.Name, "", false))
		}
	}
	page.search = buildSearchForm(components...)
}
func buildTable(permission string) string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf(`      <a  v-if="$auth('%s.edit')" @click="handleEdit(record)">编辑</a>
            <a-divider  v-if="$auth('%s.edit')" type="vertical" />
            <a v-if="$auth('%s.del')" @click="handleSub(record)">删除</a>`, permission, permission, permission))
	return buf.String()
}

func newColumn(title, name string) string {
	buf := strings.Builder{}
	buf.WriteString("{")
	buf.WriteString(fmt.Sprintf("title:'%s'", title))
	buf.WriteString(fmt.Sprintf("dataIndex:'%s'", name))
	buf.WriteString(fmt.Sprintf("key:'%s'", name))
	buf.WriteString("}")
	return buf.String()
}

func buildColumn(field []core.Field) string {
	var columns []string
	for _, c := range field {
		columns = append(columns, newColumn(c.HName, c.Name))
	}
	return strings.Join(columns, "\n")
}



func parseDom(dom core.Dom) page {
	var p page
	getSearchComponent(&p, dom)
	p.topOperator = buildTopOperator(dom.Name)
	p.table = buildTable(dom.Name)
	p.column = buildColumn(dom.Fields)

	return p
}
