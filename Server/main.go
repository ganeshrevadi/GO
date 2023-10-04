package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Server")
	GetReq()
}

func GetReq() {
	const myurl = "http://localhost:8000"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println(res.Body)
	fmt.Println("Status Code:", res.StatusCode)
}
