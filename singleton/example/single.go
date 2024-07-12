package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	// for checkign whether the instance is already created or not
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// this check is for all the goroutines that passed the above check
		if singleInstance == nil {
			fmt.Println("Creating Single Instance Now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single Instance already created.")
		}
	} else {
		fmt.Println("Single Instance already created.")
	}
	return singleInstance
}
