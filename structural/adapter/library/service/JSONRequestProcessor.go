package service

import "fmt"

type JSON struct {
	field string
}

func NewJSONRequest(field string) *JSON {
	return &JSON{field: field}
}

type JSONRequestProcessor struct {

}

func (p *JSONRequestProcessor) ProcessRequest(request *JSON) {
	fmt.Println("done processing request with request body:", request)
}

