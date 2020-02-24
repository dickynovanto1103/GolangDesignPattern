package adapter

import (
	"GolangDesignPattern/structural/adapter/library/model"
	"GolangDesignPattern/structural/adapter/library/service"
)

type AdapterXMLToJSON struct {
	service *service.JSONRequestProcessor
}

func NewAdapterXMLToJSON(service *service.JSONRequestProcessor) *AdapterXMLToJSON {
	return &AdapterXMLToJSON{service:service}
}

func (adapter *AdapterXMLToJSON) Execute(xml *model.XML) {
	jsonRequest := convertXMLToJSON(xml)
	adapter.service.ProcessRequest(jsonRequest)
}

func convertXMLToJSON(xml *model.XML) *model.JSON {
	return model.NewJSONRequest(xml.GetField())
}