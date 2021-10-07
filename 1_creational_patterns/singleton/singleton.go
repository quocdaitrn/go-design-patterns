package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct {}

var instance *singleton

// GetInstance retrieve the instance of singleton.
func GetInstanse() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		// double check locking
		if instance == nil {
			fmt.Println("Create an instance of singleton")
			instance = &singleton{}
		}
	}

	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			GetInstanse()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("DONE!")
}