package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Https Requestes")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)

	fmt.Printf("The Response type is : %T", res)

	defer res.Body.Close()

	databases, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(databases))

}
