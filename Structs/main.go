package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	Ganesh := User{"Ganesh", "ganesh@go.dev", true, 19}
	fmt.Println(Ganesh)
	fmt.Printf("The Details of User are : %+v\n", Ganesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
