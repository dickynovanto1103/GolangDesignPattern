package main

import (
	"GolangDesignPattern/structural/adapter/library/adapter"
	service2 "GolangDesignPattern/structural/adapter/library/service"
)

func main() {
	xmlRequest := adapter.NewXMLRequest("request body")

	service := &service2.JSONRequestProcessor{}

	adapterXMLToJSON := adapter.NewAdapterXMLToJSON(service)
	adapterXMLToJSON.Execute(xmlRequest)
}
