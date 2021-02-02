package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		print("first goroutine", 4, 3)
		wg.Done()
	}()

	go func() {
		go print("second goroutine", 2, 4)
		wg.Done()
	}()

	go func() {
		go print("third goroutine", 3, 5)
		wg.Done()
	}()

	wg.Wait()
}

func print(texto string, times int, wait int) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Second * time.Duration(wait))
		fmt.Println(texto)
	}
}
