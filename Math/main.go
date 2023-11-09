package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	ran := rand.Intn(10)

	fmt.Println(ran)
}
