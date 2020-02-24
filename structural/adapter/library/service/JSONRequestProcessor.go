package service

import (
	"GolangDesignPattern/structural/adapter/library/model"
	"fmt"
)

type JSONRequestProcessor struct {

}

func (p *JSONRequestProcessor) ProcessRequest(request *model.JSON) {
	fmt.Println("done processing request with request body:", request)
}

