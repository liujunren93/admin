package checkbox

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
	"strings"
)

const (
	base = "<a-checkbox value=%+v >%s</a-checkbox>"
)

type checkbox struct {
	component.Component
}

func (c *checkbox) Import() string {
	return ""
}

func (c *checkbox) Html(t string) string {
	return fmt.Sprintf(base, c.DefVal, c.Label)
}

func NewCheckbox(label string, val interface{}) component.Componenter {
	return &checkbox{component.Component{
		DefVal: val,
		Label:  label,
	}}
}

type group struct {
	checkboxs []component.Componenter
	component.Component
}

func NewGroup(name, bindModel, label string, val interface{}, required bool, checkboxs ...component.Componenter) component.Componenter {
	return &group{
		checkboxs: checkboxs,
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

func (g *group) Html(t string) string {

	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf(t+"<a-checkbox-group\n%s\tname=%q", t, g.Name))
	if g.BindModel != "" {
		buf.WriteString(fmt.Sprintf("\n%s\tv-model=%q", t, g.BindModel))
	} else {
		buf.WriteString("\n\t"+t)
		buf.WriteString(utils2.Decorator(g.Name, "请选择"+g.Label, g.DefVal, g.IsRequired))
	}
	buf.WriteString(">\n")
	for _, componenter := range g.checkboxs {
		buf.WriteString("\t\t"+t+ componenter.Html(""))
		buf.WriteString("\n")
	}
	buf.WriteString(t+"</a-checkbox-group>")
	return buf.String()
}

func (g *group) Import() string {
	return ""
}
