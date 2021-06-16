package router

import (
	"fmt"
	"strings"
)

type router struct {
	path      string
	name      string
	component string
	meta      string //{ title: 'menu.dashboard.analysis', keepAlive: false, permission: ['dashboard'] }
	redirect  string
	children  routerList
}
type routerList []router

func (r router) addChildren(router2 router) {
	r.children = append(r.children, router2)
}

func (rs routerList) String() string {
	buf := strings.Builder{}
	buf.WriteString("[\n")
	rl := len(rs)
	for i, r := range rs {
		buf.WriteString(r.String())
		if i < rl {
			buf.WriteString(",")
		}
		buf.WriteString("\n")
	}
	buf.WriteString("]")
	return buf.String()
}

func (r router) String() string {
	buf := strings.Builder{}
	buf.WriteString("{\n")
	buf.WriteString(fmt.Sprintf("\tpath:'%s',\n", r.path))
	buf.WriteString(fmt.Sprintf("\tname:'%s',\n", r.name))
	buf.WriteString(fmt.Sprintf("\tcomponent:'%s',\n", r.component))
	buf.WriteString(fmt.Sprintf("\tmeta:'%s',\n", r.meta))
	buf.WriteString(fmt.Sprintf("\tredirect:'%s',\n", r.redirect))
	if len(r.children) != 0 {
		buf.WriteString(fmt.Sprintf("\tchildren:'%s',\n", r.children.String()))
	}
	buf.WriteString("}")
	return buf.String()
}

type index struct {
	imports []string
	routers  []string
}

func (i index) add(im, router string) {
	i.imports = append(i.imports, im)
	i.routers = append(i.routers, router)
}

func (i index) String() string {
	buf := strings.Builder{}
	for _, s := range i.imports {
		buf.WriteString(s + "\n")
	}
	buf.WriteString("const routers =[\n")
	buf.WriteString(strings.Join(i.routers, ",\n"))
	buf.WriteString("]\n")
	buf.WriteString("\nexport { routers }\n")
	return buf.String()
}
