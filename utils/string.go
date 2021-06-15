package utils

import (
	"strings"
)

// 驼峰转蛇形
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	lower := strings.ToLower(string(data[:]))
	return strings.Replace(lower, "i_d", "id", -1)

}

func UcFirst(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[:1]) + str[1:]
	}
	return ""

}
func UpFirst(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[:1]) + str[1:]
	}
	return ""
}

func TagParse(str string) map[string]string {
	str=strings.Trim(str,"\r\n")
	var tagMap = map[string]string{}

	split := strings.Split(str, ";\r\n")

	for _, s := range split {
		tags := strings.Split(s, ":")
		tagMap[tags[0]] = tags[1]
	}

	return tagMap
}

func uniqueSliceStr(data []string)[]string  {
	i:=0
	var uniqueMap =make(map[string]struct{})
	for _, da := range data {
		if _,ok:=uniqueMap[da];!ok {

		}
	}
}