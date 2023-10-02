package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files !")
	content := "This needs to go in a file "

	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	datatable, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text in the file is :  \n", string(datatable))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
