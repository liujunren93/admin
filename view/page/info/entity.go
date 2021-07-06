package info

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/utils"
	"github.com/liujunren93/admin/view/base"
	"github.com/liujunren93/admin/view/component"
	"github.com/liujunren93/admin/view/component/checkbox"
	"github.com/liujunren93/admin/view/component/editor"
	"github.com/liujunren93/admin/view/component/radio"
	_select "github.com/liujunren93/admin/view/component/select"
	"github.com/liujunren93/admin/view/component/textarea"
	"github.com/liujunren93/admin/view/component/upload"
	page2 "github.com/liujunren93/admin/view/page"
	"strings"
)

type page struct {
	formItem   string
	components []string
	importList []string
	data       []string
}

func build(dom core.Dom) page {
	var p page
	//buf:=strings.Builder{}
	var components []component.Componenter
	for _, field := range dom.Fields {
		val := field.FindTagVal(core.TagBind, "required")
		isRequired := false
		if val != "" {
			isRequired = true
		}

		switch {
		case strings.Index(field.HType, "radio") >= 0:
			options, defval, isJson :=page2.MustCompile(`radio\((.*)\)`, field.HType)
			components = append(components, radio.NewGroupWithOpts(field.Name, field.HName, "", defval, isRequired, options))
			if !isJson {
				p.importList = append(p.importList, fmt.Sprintf("import { list as %sList } from '@/api/%s'", utils.UcFirst(options), utils.UpFirst(options)))
				p.data = append(p.data, fmt.Sprintf("\t"+`%sData: () => {return  %sList().then(res => {return res.data})},`,
					utils.UcFirst(options), utils.UcFirst(options)))
			}
		case strings.Index(field.HType, "checkbox") >= 0:
			options, defval, isJson := page2.MustCompile(`checkbox\((.*)\)`, field.HType)
			components = append(components, checkbox.NewGroupWithOpts(field.Name, field.HName, "", defval, isRequired, options))

			if !isJson {
				p.importList = append(p.importList, fmt.Sprintf("import { list as %sList } from '@/api/%s'", utils.UcFirst(options), utils.UpFirst(options)))
				p.data = append(p.data, fmt.Sprintf("\t"+`%sData: () => {return  %sList ().then(res => {return res.data	})	},`,
				utils.UcFirst(options), utils.UcFirst(options)))
			}
		case "textarea" == field.HType:
			components = append(components, textarea.NewTextArea(field.Name, field.HName, "", isRequired))
		case "html" == field.HType:
			newEditor := editor.NewEditor(field.Name, field.HName, "", isRequired)
			components = append(components, newEditor)
			p.importList = append(p.importList, newEditor.GetImport())
			p.components = append(p.components, newEditor.GetComponent())

		case "img" == field.HType:
			img := upload.NewUpImg(field.Name, field.HName, "", isRequired)
			components = append(components, img)
			p.importList = append(p.importList, img.GetImport())
			p.components = append(p.components, img.GetComponent())
		case strings.Index(field.HType, "select") >= 0:
			options, def, isJson := page2.MustCompile(`select\((.*)\)`, field.HType)
			components = append(components, _select.NewSimple(field.Name, field.HName, "", def, isRequired, options))

			if !isJson {
				p.importList = append(p.importList, fmt.Sprintf("import { list as %sList } from '@/api/%s'",  utils.UcFirst(options), utils.UpFirst(options)))
				p.data = append(p.data, fmt.Sprintf("\t"+`%sData: () => { return  %sList().then(res => {return res.data})},`, utils.UcFirst(options), utils.UcFirst(options)))
			}

		}
	}
	p.importList = append(p.importList,fmt.Sprintf("import { create,info,update } from '@/api/%s'",dom.Name) )

	item := buildItem(components)
	p.formItem = item
	return p

}
func buildItem(cs []component.Componenter) string {
	buf := strings.Builder{}
	var t base.T
	for _, c := range cs {
		buf.WriteString(fmt.Sprintf("\n\t\t<a-fo1rm-item\n          :label=%q\n          :labelCol=\"{lg: {span: 7}, sm: {span: 7}}\"\n          :wrapperCol=\"{lg: {span: 10}, sm: {span: 17} }\"> \n", c.GetLabel()))
		buf.WriteString(c.GetHtml(t.Multiple(2)))
		buf.WriteString("\n")
		buf.WriteString("\t\t</a-form-item>")
	}
	return buf.String()
}

func parseDom(dom core.Dom) page {
	p := build(dom)
	p.data = utils.UniqueSliceStr(p.data)
	p.importList = utils.UniqueSliceStr(p.importList)
	p.components = utils.UniqueSliceStr(p.components)
	return p
}
