package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type TodoList struct {
	TodoList []Todo `json:"todo_list"`
}

type Todo struct {
	TaskId int    `json:"taskId"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func parseTodoList() TodoList {
	jsonFile, err := os.Open("todo.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opening Stored Todo List.")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var todoList TodoList

	json.Unmarshal(byteValue, &todoList)

	return todoList

}

func printTodoList(todoList TodoList) {
	for i := 0; i < len(todoList.TodoList); i++ {
		fmt.Print("Task # : " + strconv.Itoa(todoList.TodoList[i].TaskId) + " | ")
		fmt.Print("Description:  " + todoList.TodoList[i].Task + " | ")
		fmt.Println(" Status: " + todoList.TodoList[i].Status + " | ")
	}
}

func addTask(item string) {
	var task Todo
	task.TaskId = 3
	task.Task = item
	task.Status = "incomplete"
	fmt.Println(task.Task)
}

func main() {
	addItem := flag.String("item", "task description", "a string var")
	viewList := flag.Bool("list", false, "view task list")

	flag.Parse()

	var todoList TodoList
	todoList = parseTodoList()
	if *viewList {
		printTodoList(todoList)
	}

	if isFlagPassed("item") {
		addTask(*addItem)
	}

}
