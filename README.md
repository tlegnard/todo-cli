# ToDo-CLI
## A ToDo List management CLI, written in Go.

A hands-on way for me to learn Go, and build something useful in the process.

## Usage
```bash
go build todo.go
```
If using for the first time, `./todo` will start a `todo.json` and ask for a string of input to set your first task.
## flags
```
-add string
        a string var (default "task description")
-delete int
    task ID 
-list
    view task list
-update int
    taskId 
```
