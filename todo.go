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

func parseJSON() TodoList {
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

func saveJson(todoList TodoList) {
	fmt.Println("Saving ToDo List to JSON")
	file, _ := json.MarshalIndent(todoList, "", " ")

	_ = ioutil.WriteFile("todo.json", file, 0644)

}

func printTodoList(todoList TodoList) {
	for i := 0; i < len(todoList.TodoList); i++ {
		fmt.Print("Task # : " + strconv.Itoa(todoList.TodoList[i].TaskId) + " | ")
		fmt.Print("Description:  " + todoList.TodoList[i].Task + " | ")
		fmt.Println(" Status: " + todoList.TodoList[i].Status + " | ")
	}
}

func addTask(item string, todoList TodoList) TodoList {
	//var task Todo
	// task.TaskId = getMaxId(todoList) //max id found
	// task.Task = item
	// task.Status = "incomplete"
	//fmt.Println("New Task added: \n" + strconv.Itoa(task.TaskId) + " " + task.Task + task.Status)
	todoList.TodoList = append(todoList.TodoList, Todo{getMaxId(todoList), item, "incomplete"})
	return todoList
}

func getMaxId(todoList TodoList) int {

	var maxVal int = 0
	for _, task := range todoList.TodoList {
		if task.TaskId > maxVal {
			maxVal = task.TaskId
		}
	}
	return maxVal + 1
}

func main() {
	addItem := flag.String("item", "task description", "a string var")
	viewList := flag.Bool("list", false, "view task list")

	flag.Parse()

	var todoList TodoList
	todoList = parseJSON()
	if *viewList {
		printTodoList(todoList)
	}

	if isFlagPassed("item") {
		todoList = addTask(*addItem, todoList)
		saveJson(todoList)
		printTodoList((todoList))
	}

}
