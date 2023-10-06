package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Server")
	GetReq()
	postJson()
	postForm()
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
	byteCount, _ := resString.Write(content)

	fmt.Println(resString.String())
	fmt.Println(byteCount, "bytes were written.")
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

	res, err := http.Post(myUrl, "application/json", req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	// var resString strings.Builder
	// fmt.Println(resString.String())

	// fmt.Println(res)
	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)
	fmt.Println(string(content))
	fmt.Println(byteCount)
	fmt.Println(resString.String())

	fmt.Println(res.StatusCode)
}

func postForm() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("name", "Rahul")
	data.Add("course", "golang")
	data.Add("price", "0")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)
	fmt.Println(string(content))
	fmt.Println(byteCount, "bytes were written.")
	fmt.Println(resString.String())

	fmt.Println(res.StatusCode)

}
