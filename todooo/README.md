# Todooo - Simple CLI Todo List Manager

A lightweight command-line todo list application written in Go that helps you manage your tasks efficiently.

## Features

- Add new tasks
- List all tasks with status indicators
- Mark tasks as completed
- Persistent storage using JSON
- Fast and simple CLI interface

## Installation

### Prerequisites
- Go 1.16 or higher

### Build from source
```bash
go build -o todooo
```

## Usage

### Add a new task
```bash
./todooo -add "Your task description"
```

### List all tasks
```bash
./todooo -list
```

### Mark a task as done
```bash
./todooo -done <task_number>
```
Note: Task numbers start from 1 (as displayed in the list)

### Show help
```bash
./todooo
```

## Examples

```bash
# Add some tasks
./todooo -add "Learn Go basics"
./todooo -add "Build a CLI tool"
./todooo -add "Write documentation"

# List all tasks
./todooo -list
# Output:
# 1: [ ] Learn Go basics
# 2: [ ] Build a CLI tool
# 3: [ ] Write documentation

# Mark first task as done
./todooo -done 1

# List tasks again
./todooo -list
# Output:
# 1: [x] Learn Go basics
# 2: [ ] Build a CLI tool
# 3: [ ] Write documentation
```

## Project Structure

```
todooo/
├── main.go      # Main application logic and CLI interface
├── todo.go      # Todo data structure and file operations
├── todos.json   # Data persistence file (created automatically)
└── todo         # Compiled binary
```

## Data Storage

Tasks are stored in `todos.json` in the following format:
```json
[
  {
    "task": "Task description",
    "done": false
  }
]
```

## Technical Details

- **Language**: Go
- **Storage**: JSON file (`todos.json`)
- **CLI Library**: Standard `flag` package
- **File Permissions**: 0644 for the JSON file

## Development

### Code Structure
- `Todo` struct defines the task data model
- `loadTodos()` handles reading tasks from JSON file
- `saveTodos()` handles writing tasks to JSON file
- Main function handles CLI argument parsing and command execution

### Building
```bash
go build -o todo
```

### Running Tests
```bash
go test
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source. Feel free to use and modify as needed.

## Future Enhancements

- [ ] Edit existing tasks
- [ ] Task priorities
- [ ] Due dates
- [ ] Categories/tags
- [ ] Export to different formats
