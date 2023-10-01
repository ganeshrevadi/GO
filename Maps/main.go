package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	lan := make(map[int]string)

	lan[1] = "AI"
	lan[2] = "cs"

	fmt.Println(lan)

}
