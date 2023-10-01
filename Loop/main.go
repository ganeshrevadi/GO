package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "fri", "sat"}
	fmt.Println(days)

	for _, i := range days {

		fmt.Println(i)
	}
}
