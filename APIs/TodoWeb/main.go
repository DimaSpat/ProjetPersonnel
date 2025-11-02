package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task
var nextID = 1

func main() {
	const DOMAIN = "localhost"
	const PORT = ":8080"
	const NETWORK = DOMAIN + PORT

	tasks = append(tasks, Task{ID: nextID, Title: "Finish Go API project", Done: false})
	nextID++
	tasks = append(tasks, Task{ID: nextID, Title: "Check the API in the browser", Done: true})
	nextID++

	http.HandleFunc("/tasks", tasksHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Simple Task API!  Try navigating to http://%s/tasks", NETWORK)
	})

	fmt.Printf("API Server starting!\nEndpoints:\n  GET    /tasks (View all tasks)\n  POST   /tasks (Create a new task)\n  PUT    /tasks (Update an existing task)\nOpen http://%s in your browser, or use a tool like 'curl' or Postman.\n", NETWORK)

	err := http.ListenAndServe(NETWORK, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(response)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respondJSON(w, http.StatusOK, tasks)
	case "POST":
		var newTask Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&newTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newTask.ID = nextID
		newTask.Done = true
		tasks = append(tasks, newTask)
		nextID++

		respondJSON(w, http.StatusOK, newTask)
	case "PUT":
		var updateTask Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&updateTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		found := false
		for i, t := range tasks {
			if t.ID == updateTask.ID {
				tasks[i].Title = updateTask.Title
				tasks[i].Done = updateTask.Done
				respondJSON(w, http.StatusOK, tasks[i])
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "Task ID not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}
