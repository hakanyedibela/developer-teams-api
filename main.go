package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Developer struct {
	JobDescription string `json:"job_description"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Age            int    `json:"age"`
}

var Developers []Developer

func Start(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "<h1>Run /developer Endpoint</h1>")
	if err != nil {
		return
	}
	log.Print(request.RequestURI)
	fmt.Println("Developer Endpoint: Start")
}

func returnAllDevelopers(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Developer Endpoint: Get all Developers")
	err := json.NewEncoder(writer).Encode(Developers)
	if err != nil {
		return
	}
	fmt.Println("Request: " + request.RequestURI)
}

func handleRequests() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/developer", returnAllDevelopers)

	log.Fatal(http.ListenAndServe(":18080", nil))
}

func main() {
	Developers = []Developer{
		Developer{JobDescription: "Solution Architect", FirstName: "Hakan", LastName: "Yedibela", Age: 38},
		Developer{JobDescription: "Full-Stack Developer", FirstName: "John", LastName: "Doe", Age: 22},
		{JobDescription: "DevOps Engineer", FirstName: "Peter", LastName: "Parker", Age: 23},
		{JobDescription: "Cloud-Engineer", FirstName: "Bruce", LastName: "Banner", Age: 54},
	}
	handleRequests()
}
