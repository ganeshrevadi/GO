package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
	Password string   `json:"-"`
}

func main() {
	fmt.Println("Json !!")
	jsonHandler()

}

func jsonHandler() {
	lco := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", []string{"web-dev", "js"}, "abc"},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", []string{"web-dev", "js", "react"}, "abc"},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", []string{"web-dev", "js", "angular"}, "abc"},
	}

	finalJson, err := json.MarshalIndent(lco, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
