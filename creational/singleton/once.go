package main

import (
	"fmt"
	"sync"
)

type insOnce struct {

}

var ins *insOnce

var once = &sync.Once{}

func getInsOnceInstance() *insOnce {
	if ins == nil {
		once.Do(func(){
			ins = &insOnce{}
			fmt.Println("ins created once here")
		})
	}
	fmt.Println("ins has been created before")
	return ins
}
