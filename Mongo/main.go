package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hey From Mongo DB Connecting to your database Enjoy ok")
	r := mux.NewRouter()
	fmt.Println(r)
}
