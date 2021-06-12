package upload

import (
	"fmt"
	"github.com/liujunren93/admin/view/component"
	utils2 "github.com/liujunren93/admin/view/utils"
)

const base = `<UpImg %s  placeholder='%s'/>`

type upImg struct {
	component.Component
}

func NewUpImg(name, label string, val interface{}, required bool) component.Componenter {
	return upImg{
		Component: component.Component{
			IsRequired: required,
			Name:       name,
			DefVal:     val,
			Label:      label,
		},
	}
}

func (u upImg) Import() string {
	return `import UploadImg from '@/components/common/uploadImg'`
}

func (u upImg) Html(t string) string {
	return fmt.Sprintf(base, utils2.Decorator(u.Name, "请上传"+u.Label, u.DefVal, u.IsRequired), "请上传"+u.Label)
}
