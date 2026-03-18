// package taskManager

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Task represents a single todo item
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const fileName = "tasks.json"

func main() {
	// 1. Get command line arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: task [add|list|done|delete] [content/id]")
		return
	}

	command := args[1]
	tasks := loadTasks()

	// 2. Route commands
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: Please provide a task description.")
			return
		}
		newTask := Task{
			ID:        len(tasks) + 1,
			Title:     args[2],
			Completed: false,
		}
		tasks = append(tasks, newTask)
		saveTasks(tasks)
		fmt.Printf("Added: %s\n", newTask.Title)

	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}

	case "done":
		if len(args) < 3 {
			fmt.Println("Error: Provide task ID.")
			return
		}
		id, _ := strconv.Atoi(args[2])
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
			}
		}
		saveTasks(tasks)
		fmt.Println("Task marked as done.")

	default:
		fmt.Println("Unknown command.")
	}
}

// loadTasks reads the JSON file into a Task slice
func loadTasks() []Task {
	var tasks []Task
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{} // Return empty if file doesn't exist
	}
	json.Unmarshal(data, &tasks)
	return tasks
}

// saveTasks writes the Task slice to the JSON file
func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(fileName, data, 0644)
}
