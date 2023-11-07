package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func main() {
	fmt.Println("Concurrency in Go")

	websites := []string{
		"http://github.com",
		"http://fb.com",
		"http://google.com",
	}
	// go greeter("Hello")

	// greeter("World")

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()

}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		// time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops")
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Println(endpoint, "- StatusCode: ", res.StatusCode)
	defer wg.Done()
}
