package component

type Componenter interface {
	GetImport() string
	GetComponent()string
	GetHtml(t string) string
	GetLabel()string
}

type Component struct {
	IsRequired bool
	BindModel  string
	Name       string
	Label      string
	DefVal     interface{}
}



func (c Component) GetLabel()string{
	return c.Label
}
