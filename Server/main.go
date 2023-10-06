package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Server")
	GetReq()
	postJson()
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

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)

	fmt.Println(resString.String())

	fmt.Println("string(content):", string(content))
}

func postJson() {
	const myUrl = "http://localhost:8000/post"

	req := strings.NewReader(`
		{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"
	  	}
	`)

	res, _ := http.Post(myUrl, "application/json", req)
	// var resString strings.Builder
	// fmt.Println(resString.String())

	// fmt.Println(res)
	fmt.Println(res.StatusCode)
}
