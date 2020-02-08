package main

import (
	"fmt"
	"sync"
)

type instance struct{}

var singleIns *instance

var lock = &sync.Mutex{}

func getInstance() *instance {
	if singleIns == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleIns == nil {
			fmt.Println("creating object")
			singleIns = &instance{}
		}
		fmt.Println("singleIns has been created after acquiring lock")
	}

	fmt.Println("singleIns has been created before")

	return singleIns
}
