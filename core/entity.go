package core

const (
	tagBind="binding"
)
type Field struct {
	HSort   int16
	Name    string
	Tag     []string
	Type    string
	HName   string
	HType   string //text;radio;checkbox;textarea;html;pic;file;select(name)|select({"k1":"v1","k2":"v2"} );
	HSearch string //like;select
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
