package datePicker

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	"github.com/liujunren93/admin/view/utils"
	"strings"
)

type ranger struct {
	component.Component
}

func (r ranger) GetComponent() string {
	return ""
}

func NewRanger(name, bindModel, label string, val interface{}, required bool) component.Componenter {

	return ranger{
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

func (r ranger) GetImport() string {
	return ""
}

func (r ranger) GetHtml(t string) string {
	buf := strings.Builder{}
	buf.WriteString(t)
	buf.WriteString(fmt.Sprintf(t+"<a-range-picker\n\t%sformat=\"YYYY-MM-DD HH:mm:ss\"",t))
	if r.BindModel != "" {
		buf.WriteString(fmt.Sprintf("\n\t%sv-model=%q", t, r.BindModel))
	} else {
		buf.WriteString("\n\t" + t)
		buf.WriteString(utils.Decorator(r.Name, r.Label, r.DefVal, r.IsRequired))
	}
	buf.WriteString(fmt.Sprintf("\n\t%s:placeholder=%q />", t,"['开始时间','结束时间']"))
	return buf.String()
}

type simple struct {
	component.Component
}

func (s *simple) GetComponent() string {
	return ""
}

func (s *simple) GetImport() string {
	return ""
}

func (s *simple) GetHtml(t string) string {
	buf := strings.Builder{}
	buf.WriteString(t+"<a-date-picker \n\t:show-time=\"{ format: 'HH:mm' }\"  \n\tformat=\"YYYY-MM-DD HH:mm\" \n" +
		"\t:show-time=\"{ defaultValue: moment('00:00:00', 'HH:mm:ss') }\"  ")
	if s.BindModel != "" {
		buf.WriteString(fmt.Sprintf("\n\tv-model='%s'", s.BindModel))
	} else {
		buf.WriteString(utils.Decorator(s.Name, s.Label+"不能为空", s.DefVal, s.IsRequired))

	}
	buf.WriteString(fmt.Sprintf("\n\t:placeholder=%s \n", s.Label))
	buf.WriteString("/>")
	return buf.String()
}

func NewSimple(name, bindModel, label string, val interface{}, required bool) component.Componenter {
	return &simple{
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
