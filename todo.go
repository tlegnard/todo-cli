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

func main() {
	//TODO (ha)
	//link -list flag to printTodoList func
	//write addTodo func
	//link -item to addTodo func
	todoItem := flag.String("item", "task description", "a string var")
	todo_list := flag.String("list", "list all todos", "str var with name of task list")

	flag.Parse()

	fmt.Println("word:", *todoItem)
	fmt.Println("numb:", *todo_list)

	var todoList TodoList
	todoList = parseTodoList()
	printTodoList(todoList)
}
