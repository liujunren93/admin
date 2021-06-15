package _select

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
	"strings"
)

type simple struct {
	option string
	component.Component
}

func (_ *simple) GetComponent() string {
	panic("implement me")
}


func NewSimple(name, label, bindModel string, val interface{}, required bool, option string) component.Componenter {
	return &simple{
		option: option,
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

func (_ *simple) GetImport() string {
	return ""
}

func (s *simple) GetHtml(t string) string {
	buf := strings.Builder{}
	buf.WriteString(t+"<a-select ")
	if s.BindModel != "" {
		buf.WriteString(fmt.Sprintf("v-model=%q", s.BindModel))
	} else {
		buf.WriteString(utils2.Decorator(s.Name, s.Label+"不能为空", s.DefVal, s.IsRequired))
	}
	buf.WriteString(fmt.Sprintf(" :options=%q", s.option))
	buf.WriteString(fmt.Sprintf(" placeholder=\"请选择%s\" />", s.Label))
	return buf.String()
}
