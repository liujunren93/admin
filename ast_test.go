package main

import (
	"fmt"
	"github.com/liujunren93/admin/core"
	"github.com/liujunren93/admin/server"
	"github.com/liujunren93/admin/view/page/api"
	"github.com/liujunren93/admin/view/page/info"
	"github.com/liujunren93/admin/view/page/table"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParseFile(t *testing.T) {
	set := token.NewFileSet()
	file, _ := parser.ParseFile(set, "./test/admin.go", nil, 4)
	ast.Print(set, file)

}

func TestParseFile1(t *testing.T) {
	file := core.ParseFile("./test/admin.go")
	fmt.Println(file)
}

func TestParsePath(t *testing.T) {
	path := core.ParsePath("./test")
	for _, group := range path {
		for _, dom := range group.List {
			fmt.Printf("%+v\n", dom)
		}
	}
	server.NewRoute(path...)

}
func TestTable(t *testing.T) {

	path := core.ParsePath("./test")

	table.BuildPage(path)

}
func TestApi(t *testing.T) {

	path := core.ParsePath("./test")

	api.BuildPage(path)

}

func TestInfo(t *testing.T) {

	path := core.ParsePath("./test")

	info.BuildPage(path)

}

func TestRbac(t *testing.T) {
	path := core.ParsePath("./test")
	server.NewRbac(path...)
	//Info.BuildInfo("./tt",path)

}
func TestConf(t *testing.T) {
	//path := core.ParsePath("./test")
	server.NewConfigFile()
	//Info.BuildInfo("./tt",path)

}
