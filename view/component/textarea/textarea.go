package textarea

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	"github.com/liujunren93/admin/view/utils"
)

type textarea struct {
	component.Component
}


func (te textarea) GetComponent() string {
	panic("implement me")
}

func (te textarea) GetImport() string {
	return ""
}

func (te textarea) GetHtml(t string) string {

	return fmt.Sprintf(t+"\t<a-textarea  :rows=\"4\" %s placeholder=%q", utils.Decorator(te.Name, "请输入"+te.Label, "", te.IsRequired))

}

func NewTextArea(name, label string, val interface{}, required bool) component.Componenter {
	return &textarea{
		Component: component.Component{
			IsRequired: required,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}
