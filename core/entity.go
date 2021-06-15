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
	Name        string
	HName       string
	HPagination bool
	HasSearch   bool
	Fields      []Field
}

type Group struct {
	Name string // filename
	List []*Dom
}
