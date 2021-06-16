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

func (c *checkbox) GetComponent() string {
	return ""
}

func (c *checkbox) GetImport() string {
	return ""
}

func (c *checkbox) GetHtml(t string) string {
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

func (g *group) GetComponent() string {
	return ""
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

func (g *group) GetHtml(t string) string {

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
		buf.WriteString("\t\t"+t+ componenter.GetHtml(""))
		buf.WriteString("\n")
	}
	buf.WriteString(t+"</a-checkbox-group>")
	return buf.String()
}

func (g *group) GetImport() string {
	return ""
}

type groupWithOpts struct {
	option string
	component.Component
}

func (g groupWithOpts) GetComponent() string {
	return ""
}

func (g groupWithOpts) GetImport() string {
	return ""
}

func (g groupWithOpts) GetHtml(t string) string {

	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf(t+"<a-checkbox-group\n%s\tname=%q", t, g.Name))
	if g.BindModel != "" {
		buf.WriteString(fmt.Sprintf("\n%s\tv-model=%q", t, g.BindModel))
	} else {
		buf.WriteString("\n\t"+t)
		buf.WriteString(utils2.Decorator(g.Name, "请选择"+g.Label, g.DefVal, g.IsRequired))
	}
	buf.WriteString(fmt.Sprintf(" :options=\"%s\"", g.option))
	buf.WriteString("/>\n")

	return buf.String()
}

func NewGroupWithOpts(name, label, bindModel string, val interface{}, required bool,opt string) component.Componenter {
	return &groupWithOpts{
		option: opt,
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
