package adapter

import "GolangDesignPattern/structural/adapter/library/model"

type ClientInterface interface {
	Execute(xml *model.XML)
}
