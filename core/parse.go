package core

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
)

func ParsePath(dir string) []*Group {
	fset := token.NewFileSet()
	abs, _ := filepath.Abs(dir)

	dirs, err := parser.ParseDir(fset, abs, func(info fs.FileInfo) bool {
		if filepath.Ext(info.Name()) != ".go" {
			return false
		}
		return true
	}, 4)
	if len(dirs) < 1 {
		return nil
	}
	if err != nil {
		panic(err)
	}
	var groupList []*Group
	for _, file := range dirs {
		for k, v := range file.Files {
			group := parse(v)
			base := filepath.Base(k)
			ext := filepath.Ext(base)
			group.Name = strings.Trim(base, ext)
			groupList = append(groupList, group)
		}
	}

	return groupList
}

func ParseFile(filename string) *Group {
	set := token.NewFileSet()
	file, _ := parser.ParseFile(set, filename, nil, 4)
	group := parse(file)
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	group.Name = strings.Trim(base, ext)
	return group
}
func parse(asf *ast.File) *Group {
	var g Group
	for _, decl := range asf.Decls {
		switch t := decl.(type) {
		case *ast.GenDecl:
			d := parseGenDecl(t)
			if d != nil {
				g.List = append(g.List, d)
			}
		case *ast.FuncDecl:

		}
	}

	return &g

}

func parseGenDecl(g *ast.GenDecl) *Dom {

	if g.Doc == nil {
		return nil
	}
	comment := make(map[string]string)
	parseComment(g.Doc.List, comment)
	if len(comment) == 0 {
		return nil
	}
	var dom Dom
	switch sp := g.Specs[0].(type) {
	case *ast.TypeSpec:
		dom = Dom{
			Name:  sp.Name.String(),
			HName: sp.Name.String(),
		}
		if v, ok := comment["name"]; ok {
			dom.HName = v
		}
		if _, ok := comment["page"]; ok {
			dom.HPagination = true
		}
		if st, ok := sp.Type.(*ast.StructType); ok {
			dom.Fields, dom.HasSearch = parseField(st)
		}

	case *ast.ImportSpec:
	}
	return &dom

}

func parseField(st *ast.StructType) (fds []Field, search bool) {
	for _, field := range st.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		var fd Field
		var comment = make(map[string]string)
		if field.Doc != nil {
			parseComment(field.Doc.List, comment)
		}
		if field.Comment != nil {
			parseComment(field.Comment.List, comment)
		}
		if val, ok := comment["search"]; ok {
			search = true
			if val == "" {
				fd.HSearch = "true"
			} else {
				fd.HSearch = val
			}

		}

		fd.HName = comment["name"]
		fd.HType = comment["type"]
		if v, ok := comment["sort"]; ok {
			parseInt, err := strconv.ParseInt(v, 10, 16)
			if err != nil {
				panic(err)
			}
			fd.HSort = int16(parseInt)

		}

		fd.Name = field.Names[0].Name
		tag := ParseTag(field.Tag.Value)
		fd.Tag = append(fd.Tag, tag[tagBind])
		if ident, ok := field.Type.(*ast.Ident); ok {
			fd.Type = ident.Name
		}
		fds = append(fds, fd)

	}
	return
}

// 读取注释
func parseComment(comment []*ast.Comment, resComment map[string]string) {
	for _, com := range comment {
		index := strings.Index(com.Text, "@curd")
		list := strings.Split(com.Text[index+6:], ";")
		for _, li := range list {
			split := strings.Split(li, "=")
			if len(split) == 2 {
				resComment[split[0]] = split[1]
			} else {
				resComment[split[0]] = ""
			}
		}

	}
}

// 读取tag
func ParseTag(tag string) map[string]string {
	setting := map[string]string{}
	tags := strings.Split(tag, ";")
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToLower(v[0]))
		if len(v) >= 2 {
			setting[k] = strings.Join(v[1:], ":")
		} else {
			setting[k] = k
		}
	}
	return setting
}
