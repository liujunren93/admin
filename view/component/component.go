package component

type Componenter interface {
	Import() string
	Html(t string) string
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
