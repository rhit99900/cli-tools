package main

import (
	"flag"
	"fmt"
	"os"
)

const fileName = "todos.json"

func main() {
	add := flag.String("add", "", "Task to add")
	list := flag.Bool("list", false, "List all todos")
	done := flag.Int("done", -1, "Mark todo as done (by index)")

	flag.Parse()

	todos, err := loadTodos(fileName)

	if(err != nil) {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	switch {
		case *add != "":
			todos = append(todos, Todo{Task: *add, Done: false})
			saveTodos(fileName, todos)
			fmt.Println("Added:", *add)

		case *list:
			for i,t := range todos {
				status := "[ ]"
				if t.Done {
					status = "[x]"					
				}
				fmt.Printf("%d: %s %s\n", i+1, status, t.Task)
			}
		
		case *done >= 0:
			index := *done - 1
			if index < 0 || index >= len(todos) {
				fmt.Println("Invalid todo index")
				return
			}
			todos[index].Done = true
			saveTodos(fileName, todos)
			fmt.Println("Marked as done:", todos[index].Task)

		default:
			fmt.Println("Usage:")
			flag.PrintDefaults()
	}
}