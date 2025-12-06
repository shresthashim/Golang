package main

import (
	"fmt"
	"sync"
)

func test(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("This is a test goroutine")
}

func main () {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test(&wg)
	}

	wg.Wait()
	
}
