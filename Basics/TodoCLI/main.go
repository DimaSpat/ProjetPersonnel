package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Title     string
	Completed bool
}

var tasks []Task

func main() {
	var indexInput int
	var taskInput, newTaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add a new task")
	fmt.Println("2. List all tasks")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Edit a task")
	fmt.Println("5. Delete a task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice from Options: ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			for {
				fmt.Print("Enter task title: ")
				scanner.Scan()
				taskInput = scanner.Text()
				if taskInput == "" {
					fmt.Println("Invalid input, no title specified, try again")
				} else {
					break
				}
			}
			taskAdd(taskInput)
		case 2:
			taskList()
		case 3:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			taskCompleted(indexInput)
		case 4:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter task title: ")
			scanner.Scan()
			newTaskInput = scanner.Text()
			taskEdit(indexInput, newTaskInput)
		case 5:
			fmt.Print("Enter index: ")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			taskDelete(indexInput)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func taskAdd(title string) {
	newTask := Task{Title: title, Completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("New task added")
}

func taskList() {
	for i, task := range tasks {
		status := "n"
		if task.Completed {
			status = "y"
		}

		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

func taskCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Completed = true
		fmt.Println("Task completed")
	} else {
		fmt.Println("Invalid task index")
	}
}

func taskEdit(index int, newTitle string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Title = newTitle
		fmt.Println("Task edited to ", newTitle)
	} else {
		fmt.Println("Invalid task index")
	}
}

func taskDelete(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task deleted")
	} else {
		fmt.Println("Invalid task index")
	}
}
