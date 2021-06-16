package page

import (
	"encoding/json"
	"regexp"
)

type option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

//@return data,default,isJson
func MustCompile(regStr, find string) (string, string, bool) {
	vals := regexp.MustCompile(regStr).FindStringSubmatch(find)
	if len(vals) >= 2 {
		if json.Valid([]byte(vals[1])) {
			var options []option
			err := json.Unmarshal([]byte(vals[1]), &options)
			if err != nil {
				panic(err)
			}
			return vals[1], options[0].Value, true
		}
		return vals[1], "", false
	}
	return "", "", false
}
