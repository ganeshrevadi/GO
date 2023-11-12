package main

import "fmt"

func main() {
	fmt.Println("Url Shortner")
  mux := httprouter.New()

  router.GET("/" ,Index)
}
