package input

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
	"strings"
)



type input struct {
	_type string
	component.Component
}

func (i *input) GetComponent() string {
	return ""
}



func NewInput(name, label,bindModel string, val interface{}, required bool) component.Componenter {
	return &input{
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
func NewInputNumber(name, label ,bindModel string, val interface{}, required bool) component.Componenter {
	return &input{
		_type: "number",
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
func (i *input) GetImport() string {
	return ""
}
func (i *input) GetHtml(t string) string {
	buf:=strings.Builder{}
	if i._type == "number" {
		buf.WriteString(t+"<a-input-number")
	}else{
		buf.WriteString(t+"<a-input")
	}

	if i.BindModel!="" {
		buf.WriteString(fmt.Sprintf(" v-model=%q",i.Name))
	}else{
		buf.WriteString(utils2.Decorator(i.Name, i.Label+"不能为空", i.DefVal, i.IsRequired))
	}
	buf.WriteString(fmt.Sprintf("\nplaceholder='请输入%s'",i.Label))
	buf.WriteString("/>")
	return buf.String()
}
