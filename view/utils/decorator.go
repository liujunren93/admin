package utils

import (
	"fmt"
	"github.com/liujunren93/admin/utils"
	"time"
)

func Decorator(name, msg string, defVal interface{}, isRequired bool) string {
	switch dv := defVal.(type) {
	case string:
		defVal = fmt.Sprintf("'%s'", dv)
	case []time.Time:
		defVal=fmt.Sprintf("['%s','%s']",dv[0],dv[1])
	case time.Time:
		defVal = fmt.Sprintf("'%s'", dv)
	}

	var rule = ""
	if isRequired {
		rule = fmt.Sprintf(`rules: [{ required: true, message: '%s' }]`, msg)
	}
	if rule != "" {
		return fmt.Sprintf(`v-decorator="['%s',{'%s',initialValue:%v}]"`, utils.SnakeString(name), rule, defVal)
	} else {
		return fmt.Sprintf(`v-decorator="['%s', {initialValue:%v }]"`,utils.SnakeString(name) , defVal)
	}

}
