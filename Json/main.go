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
	DecodeJson()

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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v", myOnlineData)
}
