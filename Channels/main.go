package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels ")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {

		val, isChannelopen := <-myCh

		fmt.Println(isChannelopen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {

		myCh <- 5
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
