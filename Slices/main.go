package main

import "fmt"

func main() {
	fmt.Println("Hello Slices !")

	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

}
