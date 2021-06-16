package radio

import (
	_ "embed"
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
	"strings"
)

const (
	baseHtml  = "<a-radio  :value='%v' >%s</a-radio>"
	groupHtml = "<a-radio-group name='%s' %s > %s </a-radio-group>"
)

type radio struct {
	component.Component
}

func (r *radio) GetComponent() string {
	return ""
}

func (r *radio) GetImport() string {
	return ""
}

func (r *radio) GetHtml(t string) string {
	return fmt.Sprintf(t+baseHtml, r.DefVal, r.Label)
}

func NewRadio(name, label, bindModel string, val interface{}, required bool) component.Componenter {
	return &radio{
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

type group struct {
	radios []radio
	component.Component
}

func NewGroup(name, label, bindModel string, val interface{}, required bool, radios ...radio) component.Componenter {
	return &group{
		radios: radios,
		Component: component.Component{
			IsRequired: required,
			BindModel:  bindModel,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
func (g *group) GetComponent() string {
	return ""
}
func (g *group) GetImport() string {
	return ""
}

func (g *group) GetHtml(t string) string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf(t+"<a-radio-group name='%s'\n", g.Name))
	if g.BindModel != "" {
		buf.WriteString(fmt.Sprintf(" v-model=\"%s\"", g.BindModel))
	} else {
		buf.WriteString(utils2.Decorator(g.Name, "请选择"+g.Label, g.DefVal, g.IsRequired))
	}
	buf.WriteString(" >")

	for _, r := range g.radios {
		buf.WriteString("\n\t\t")
		buf.WriteString(r.GetHtml(""))
	}
	buf.WriteString("\n")
	buf.WriteString(t+"</a-radio-group>")
	return buf.String()
}

type groupWithOpts struct {
	option string
	component.Component
}

func (g groupWithOpts) GetComponent() string {
	return ""
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

func (g groupWithOpts) GetImport() string {
	panic("implement me")
}

func (g groupWithOpts) GetHtml(t string) string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf(t+"<a-radio-group name='%s'\n", g.Name))
	if g.BindModel != "" {
		buf.WriteString(fmt.Sprintf(" v-model=%q", g.BindModel))
	} else {
		buf.WriteString(utils2.Decorator(g.Name, "请选择"+g.Label, g.DefVal, g.IsRequired))
	}
	buf.WriteString(fmt.Sprintf(" :options=\"%s\"", g.option))
	buf.WriteString("/>")



	return buf.String()
}
