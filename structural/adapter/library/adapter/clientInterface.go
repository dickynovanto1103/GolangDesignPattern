package adapter

type XML struct {
	field string
}

func NewXMLRequest(field string) *XML {
	return &XML{field:field}
}

type ClientInterface interface {
	Execute(xml XML)
}
