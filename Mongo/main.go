package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hey From Mongo!")
	r := mux.NewRouter()
	fmt.Println(r)
}
