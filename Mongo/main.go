package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("")
	r := mux.NewRouter()
	fmt.Println(r)
}
