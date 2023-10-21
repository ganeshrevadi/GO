package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"`
}

type Author struct {
	FulName string `json:"full_name"`
	Website string `json:"website"`
}

var courses []Course

// MiddleWare , helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Build API with Golang")
	router := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseID: "1", CourseName: "Golang", CoursePrice: 100, Author: &Author{FulName: "Golang", Website: "golang.org"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "Python", CoursePrice: 200, Author: &Author{FulName: "Python", Website: "python.org"}})

	router.HandleFunc("/", serverHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourse).Methods("GET")
	router.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourseOne).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))

}

//controllers - file

// serve home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in GOLANG FROM VS</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Course List \n"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one Course !")
	w.Header().Set("Content-type", "application/json")

	//grab id from request

	params := mux.Vars(r)

	//loop through course and find matching id and return response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course Found !")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	//what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data")
		return
	}

	if checkForDuplicateCourseID(course.CourseID, courses) {
		json.NewEncoder(w).Encode("Data already exist !")
		return
	}

	//generate unique id
	// append course into course
	rand.Seed(time.Now().UnixNano())

	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourseOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop  , id , remove  , add with my ID

	for i, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			break
		}
	}
}

func checkForDuplicateCourseID(courseID string, courses []Course) bool {
	for _, course := range courses {
		if course.CourseID == courseID {
			return true
		}
	}
	return false
}
