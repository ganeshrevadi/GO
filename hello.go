package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World")

	fmt.Println("Enter the Rating:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Print("Thanks for the Rating: ", input, " stars")

}
