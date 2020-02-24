package model

type XML struct {
	field string
}

func NewXMLRequest(field string) *XML {
	return &XML{field:field}
}

func (x *XML) GetField() string {
	return x.field
}