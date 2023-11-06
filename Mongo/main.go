package main

import (
	"fmt"
	"ganeshrevadi/GO/Mongo/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hey From Mongo DB Connecting to your database !")
	r := routes.Router()
	fmt.Println("Server is getting Started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing On Port 4000 ...")
}
