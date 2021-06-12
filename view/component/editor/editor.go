package editor

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
)

const base = `<Editor %s  placeholder='%s'/>`

type editor struct {
	component.Component
}

func NewEditor(name, label string, val interface{}, required bool) component.Componenter {
	return &editor{
		Component: component.Component{
			IsRequired: required,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

func (e *editor) Html(t string) string {
	return fmt.Sprintf(t+base, utils2.Decorator(e.Name, e.Label+"不能为空", e.DefVal, e.IsRequired),  "请输入"+e.Label)
}

func (e *editor) Import() string {
	return `import Editor from '@/components/Editor/tinymce/Tinymce'`
}
