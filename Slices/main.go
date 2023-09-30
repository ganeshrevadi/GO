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

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
}
