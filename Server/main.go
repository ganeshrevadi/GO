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
	fmt.Println(res.Status)

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)

	fmt.Println(resString.String())

	fmt.Println("string(content):", string(content))
}
