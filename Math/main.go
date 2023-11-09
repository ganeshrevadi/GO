package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// rand.NewSource(time.Now().UnixNano())
	// ran := rand.Intn(10)

	// fmt.Println(ran)
	myRan, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRan)
}
