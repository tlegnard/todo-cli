# ToDo-CLI
## A ToDo List management CLI, written in Go.

A hands-on way for me to learn Go, and build something useful in the process.

TODO (ha)
- [X] add command line flags
- [X] read JSON and print todos
- [X] split read/print into separate functions
- [X] link -list flag to printTodoList func
- [X] write addTodo func
  - [X] add default status "incomplete" and increment index of TaskId for new tasks
- [X] link -item to addTodo func
- [X] add delete item flag and function
- [X] add update task status item
  - [make sure \n doesn't get saved to json]
- [ ] migrate json to sqlite or better storage method
  - [ ] add schema
- [ ] add todolist initialization (goes by default if no file found)
- [ ] Update this readme to provide usage
- [ ] add go module
  - [ ] move separate functions to new file, saved in go mod