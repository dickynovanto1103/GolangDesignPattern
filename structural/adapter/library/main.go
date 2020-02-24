package main

import (
	"GolangDesignPattern/structural/adapter/library/adapter"
	"GolangDesignPattern/structural/adapter/library/model"
	service2 "GolangDesignPattern/structural/adapter/library/service"
)

func main() {
	xmlRequest := model.NewXMLRequest("request body")

	service := &service2.JSONRequestProcessor{}

	adapterXMLToJSON := adapter.NewAdapterXMLToJSON(service)
	adapterXMLToJSON.Execute(xmlRequest)
}
