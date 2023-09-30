package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello Slices !")

	var fruitList = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	highScores := make([]int, 2)

	highScores[0] = 23
	highScores[1] = 56

	fmt.Println(highScores)

	highScores = append(highScores, 43, 342)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//Removing Values from Slices based on index

	var index = 2
	highScores = append(highScores[:index], highScores[index+1:]...)

	fmt.Println(highScores)
}
