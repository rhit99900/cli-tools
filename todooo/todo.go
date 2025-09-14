package main

import (
	"encoding/json"
	"os"	
)

type Todo struct {
	Task	string	`json:"task"`
	Done	bool		`json:"done"`
}


func loadTodos(filename string) ([]Todo, error) {
	var todos []Todo
	file, err := os.ReadFile(filename)

	if err!= nil {
		if os.IsNotExist(err) {
			return todos, nil // Return empty list if file doesn't exist.			
		}
		return nil, err		
	}
	err = json.Unmarshal(file, &todos)
	return todos, err
}

func saveTodos(filename string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}