package Info

import (
	"encoding/json"
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
	"regexp"
	"strings"
)

type page struct {
	formItem   string
	components []string
	importList []string
	data       []string
}

func build( dom core.Dom) page {
	var p  page
	//buf:=strings.Builder{}
	var components []component.Componenter
	for _, field := range dom.Fields {
		val := field.FindTagVal(core.TagBind, "required")
		isRequired := false
		if val != "" {
			isRequired = true
		}
		upName := utils.UpFirst(field.Name)
		ucName := utils.UcFirst(field.Name)
		switch {
		case strings.Index(field.HType, "radio") >= 0:
			if options, defval, isJson := mustCompile("radio\\((\\w*)\\)", field.HType); isJson {
				components = append(components, radio.NewGroupWithOpts(field.Name, field.HName, "", defval, isRequired, options))
			} else {
				p.importList = append(p.importList, fmt.Sprintf("import { list as get%sList } from '@/api/%s'", upName, upName))
				p.data = append(p.data, fmt.Sprintf(`              %sData: () => {
                      return  get%sList ()
                        .then(res => {
                          return res.data
                        })
                    },`, upName, ucName))
			}
		case strings.Index(field.HType, "checkbox") >= 0:
			if options, defval, isJson := mustCompile("radio\\((\\w*)\\)", field.HType); isJson {
				components = append(components, checkbox.NewGroupWithOpts(field.Name, field.HName, "", defval, isRequired, options))
			} else {
				p.importList = append(p.importList, fmt.Sprintf("import { list as get%sList } from '@/api/%s'", upName, upName))
				p.data = append(p.data, fmt.Sprintf(`              %sData: () => {
                      return  get%sList ()
                        .then(res => {
                          return res.data
                        })
                    },`, upName, ucName))
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
			if compile, def, isJson := mustCompile(`select\((\w*)\)`, field.HType); isJson {
				components = append(components, _select.NewSimple(field.Name, field.HName, "", def, isRequired, compile))
			} else {
				p.importList = append(p.importList, fmt.Sprintf("import { list as %sList } from '@/api/%s'", ucName, upName))
				p.data = append(p.data, fmt.Sprintf(`              %sData: () => {
                      return  %slist ()
                        .then(res => {
                          return res.data
                        })
                    },`, ucName, ucName))
			}

		}

	}
	item := buildItem(components)
	p.formItem=item
	return p

}
func buildItem(cs []component.Componenter) string {
	buf:=strings.Builder{}
	var t base.T
	for _, c := range cs {
		buf.WriteString(fmt.Sprintf("    <a-form-item\n          :label=%q\n          :labelCol=\"{lg: {span: 7}, sm: {span: 7}}\"\n          :wrapperCol=\"{lg: {span: 10}, sm: {span: 17} }\">",c.GetLabel()))
		buf.WriteString(c.GetHtml(t.Multiple(1)))
		buf.WriteString(" \n</a-form-item>")
	}
	return buf.String()
}

//@return data,default,isJson
func mustCompile(regStr, find string) (string, string, bool) {
	vals := regexp.MustCompile(regStr).FindStringSubmatch(find)
	if len(vals) > 2 {

		if json.Valid([]byte(vals[1])) {
			defval := ""
			split := strings.Split(vals[1], ",")
			for _, s := range split {
				data := strings.Split(s, ":")
				defval = data[0]
				break
			}
			return vals[1], defval, true
		}
	} else {
		return vals[1], "", false
	}
	return "", "", false
}

func parseDom(dom core.Dom) page {
	p := build(dom)
	var cj int
	var coMap map[string]struct{}
	for _, s := range p.components {
		if _,ok:=coMap[s];ok {
			p.components[cj]=s
			cj++
		}else{
			
		}
	}
}

func ()  {
	
}