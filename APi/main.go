package main

import "fmt"

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FulName string `json:"full_name"`
	Website string `json:"website"`
}

var courses []Course

func main() {
	fmt.Println("Build API with Golang")
}
