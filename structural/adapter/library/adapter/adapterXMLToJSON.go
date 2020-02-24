package adapter

import (
	"GolangDesignPattern/structural/adapter/library/service"
)

type AdapterXMLToJSON struct {
	service *service.JSONRequestProcessor
}

func NewAdapterXMLToJSON(service *service.JSONRequestProcessor) *AdapterXMLToJSON {
	return &AdapterXMLToJSON{service:service}
}

func (adapter *AdapterXMLToJSON) Execute(xml *XML) {
	jsonRequest := convertXMLToJSON(xml)
	adapter.service.ProcessRequest(jsonRequest)
}

func convertXMLToJSON(xml *XML) *service.JSON {
	return service.NewJSONRequest(xml.field)
}