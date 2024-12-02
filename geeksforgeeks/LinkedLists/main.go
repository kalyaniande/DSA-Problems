package main

import (
	"fmt"
	"sync"
)

var done = make(chan bool)
var ch = make(chan int)
var wg sync.WaitGroup

func producer() {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	wg.Done()
}

func consumer() {
	for v := range ch {
		fmt.Println(v)
	}
	done <- true
}

func main() {

	wg.Add(1)
	go producer()

	wg.Add(1)
	go producer()

	go func() {
		wg.Wait()
		close(ch)
	}()

	go consumer()

	<-done
}
