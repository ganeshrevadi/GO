package main

import (
	"fmt"
	"ganeshrevadi/GO/Mongo/controller"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hey From Mongo DB Connecting to your database !")
	r := mux.NewRouter()
	fmt.Println(r)

	controller.Init()

}
