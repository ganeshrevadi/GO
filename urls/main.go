package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?name=reactjs&pay=fljalk"

func main() {
	fmt.Println("Welcome to handling URLS in GOLANG")
	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["pay"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}

	an := partsOfUrl.String()
	fmt.Println(an)
	fmt.Println("Done for the day !")
}
