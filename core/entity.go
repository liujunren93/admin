package core

const (
	TagBind = "binding"
)

type Field struct {
	HSort   int16
	Name    string
	Tag     []string
	Type    string
	HName   string
	HType   string //text;radio([{"label":"label","value":"value"}]);checkbox([{"label":"label","value":"value"}]);textarea;html;img;select(name)|select({"k1":"v1","k2":"v2"} );
	HSearch string //like;select
}

func (f *Field) FindTagVal(tag, attribute string) string {
	for _, s := range f.Tag {
		if tag == s {
			parseTag := ParseTag(s)
			return parseTag[attribute]

		}
	}
	return ""
}

type Dom struct {
	Name        string //struct name
	HName       string //网页名
	HPagination bool // 是否分页
	HasSearch   bool // 是否支持搜索
	Fields      []Field
}

type Group struct {
	Name string // filename
	List []*Dom
}
