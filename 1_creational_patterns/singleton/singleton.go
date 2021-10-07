package main

import (
	"fmt"
	"sync"
	"time"
)


type singleton struct {}

var instance *singleton // not volatile ==> double-checked locking may be broken

// var lock = &sync.Mutex{}

// // GetInstance retrieve the instance of singleton.
// func GetInstanse() *singleton {
// 	if instance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		// double check locking
// 		if instance == nil {
// 			fmt.Println("Create an instance of singleton")
// 			instance = &singleton{}
// 		}
// 	}

// 	return instance
// }

var once sync.Once

func GetInstance() *singleton {
  once.Do(func(){
    time.Sleep(5 * time.Second)
	fmt.Println("Create an instance of singleton")
    instance = &singleton{}
  })
  return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			GetInstance()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("DONE!")
}