package server

import (
	"fmt"
	"strings"
)

type File struct {
	Pkg      string // ctrl,entity,dao,router
	Name     string
	Class    []Class
	Import   []string
	Variable []string
	Funcs    []Func
}

type Class struct {
	Name   string
	Fields []Field
	Funcs  []Func
}

type Field struct {
	Name string
	Type string
	Tag  string
}
type Func struct {
	Name  string
	Param string
	Ret   string
	Body  []string
}

func (f *File) AddFunc(fun Func) *File {
	f.Funcs = append(f.Funcs, fun)
	return f
}
func (f *File) AddVariable(v string) *File {
	f.Variable = append(f.Variable, v)
	return f
}
func (f *File) String() string {
	var fBuf strings.Builder
	fBuf.WriteString(fmt.Sprintf("package %s", f.Pkg))
	if len(f.Import) > 0 {
		fBuf.WriteString("\n import ( \n")
		for _, s := range f.Import {
			fBuf.WriteString(fmt.Sprintf("\t\"%s\"\n", s))
		}
		fBuf.WriteString(")\n")
	}
	if len(f.Variable)>0 {
		for _, s := range f.Variable {
			fBuf.WriteString(s+"\n")
		}
	}
	if len(f.Funcs)>0 {
		for _, fc := range f.Funcs {
			fBuf.WriteString(fc.String("","")+"\n")
		}
	}

	for _, class := range f.Class {
		fBuf.WriteString(class.String(f.Pkg))
	}
	return fBuf.String()
}
func (c *Class) String(fileType string) string {
	var classBody strings.Builder
	for _, field := range c.Fields {
		classBody.WriteString(fmt.Sprintf("\t%s %s `%s`\n", field.Name, field.Type, field.Tag))
	}
	var funcBuf strings.Builder
	for _, f := range c.Funcs {
		funcBuf.WriteString("\n")
		funcBuf.WriteString(f.String(fileType, c.Name))

	}
	return fmt.Sprintf("type %s struct{\n%s} \n %s", c.Name, classBody.String(), funcBuf.String())
}
func (c *Class) AddFunc(f Func) *Class {
	c.Funcs = append(c.Funcs, f)
	return c
}

//NewFunc
func NewFunc(name, parameter, ret, body string) Func {
	return Func{
		Name:  name,
		Param: parameter,
		Ret:   ret,
		Body:  []string{body},
	}

}
func (f *Func) String(classType, className string) string {
	var funcBuf strings.Builder

	ret:=f.Ret
	if len(strings.Split(f.Ret, " "))>1 {
		ret="("+ret+")"
	}
	if className != "" {
		funcBuf.WriteString(fmt.Sprintf("func(*%s)%s(%s) %s {\n",  className, f.Name, f.Param, ret))
	} else {
		funcBuf.WriteString(fmt.Sprintf("func %s(%s)%s{ \n", f.Name, f.Param, ret))
	}
	for _, s := range f.Body {
		funcBuf.WriteString("\t")
		funcBuf.WriteString(s)
	}
	funcBuf.WriteString("\n}")
	return funcBuf.String()
}

func (c *Class) AddField(name, fType string, tags ...string)  {
	var tagBuf strings.Builder
	for _, s := range tags {
		tagBuf.WriteString(s + " ")
	}
	c.Fields = append(c.Fields, Field{
		Name: name,
		Type: fType,
		Tag:  tagBuf.String(),
	})

}
