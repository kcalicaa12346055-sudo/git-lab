package main

import (
	"fmt"
)

// Task interface defines the behavior for any task type
type TaskManager interface {
	Display() string
}

// Task struct holds data for a single task
type Task struct {
	ID        int
	Title     string
	IsDone    bool
}

// Display implements the TaskManager interface for our Task struct
func (t Task) Display() string {
	status := " [ ] "
	if t.IsDone {
		status = " [x] "
	}
	return fmt.Sprintf("%d.%s %s", t.ID, status, t.Title)
}

func main() {
	// Creating a slice (dynamic array) of Tasks
	todoList := []Task{
		{ID: 1, Title: "Install Go Compiler", IsDone: true},
		{ID: 2, Title: "Finish PROLANG Activity", IsDone: false},
		{ID: 3, Title: "Push code to GitHub", IsDone: false},
	}

	fmt.Println("=== GOPHER TASK MANAGER ===")
	
	// Looping through the list using the interface
	for _, item := range todoList {
		printTask(item)
	}
}

// This function accepts the interface, allowing for polymorphism
func printTask(tm TaskManager) {
	fmt.Println(tm.Display())
}