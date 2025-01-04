package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var fullGolang string = "Watch Nana's Go lang Full Course"
var shortGolang string = "Watch Go lang Crash Course"
var rewardDessert string = "Reward myself with a cheesecake"
var taskItems []string = []string{fullGolang, shortGolang, rewardDessert}

func main() {
	fmt.Print("###### Todo list API ######\n")
	http.HandleFunc("/", helloUser)         // curl http://localhost:8080/
	http.HandleFunc("/show-task", showTask) // curl http://localhost:8080/show-task
	http.HandleFunc("/add-task", addTask)   // curl -X POST -H "Content-Type: application/json" -d "\"Clear the car\"" http://localhost:8080/add-task
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greetings string = "Hello User! Welcome to our TODO list app!"
	fmt.Fprintln(writer, greetings)
}

func showTask(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var newTask string
	if err := json.Unmarshal(body, &newTask); err != nil {
		http.Error(writer, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	taskItems = append(taskItems, newTask)
	fmt.Fprintln(writer, "Task added successfully!")
}
