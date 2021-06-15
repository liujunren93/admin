package api

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/utils"

	"os"
)

func BuildApi(path string, groups []*core.Group) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		panic(err)
	}
	for _, group := range groups {
		buildFile(path, group)
	}
}
func buildFile(path string, g *core.Group) {

	for _, dom := range g.List {
		create, err := os.Create(path + "/" + dom.Name + ".js")
		if err != nil {
			panic(err)
		}
		create.WriteString("import * as http from '@/utils/http'\n")
		create.WriteString(fmt.Sprintf("const api='%s'\n", utils.UcFirst(dom.Name)))
		create.WriteString(fmt.Sprintf("export function list (parameter) {\n   return http.get(api, parameter)\n}\n"))
		create.WriteString(fmt.Sprintf("export function create (parameter) {\n   return http.post(api, parameter)\n}\n"))
		create.WriteString(fmt.Sprintf("export function info (id,parameter) {\n   return http.get(api+'/'+id, parameter)\n}\n"))
		create.WriteString(fmt.Sprintf("export function update (id,parameter) {\n   return http.put(api+'/'+id, parameter)\n}\n"))
		create.WriteString(fmt.Sprintf("export function del (id,parameter) {\n   return http.delete(api+'/'+id, parameter)\n}\n"))
	}
}
