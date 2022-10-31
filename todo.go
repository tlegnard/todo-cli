package main

import (
	"bufio"
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

func parseJSON() (TodoList, error) {
	jsonFile, err := os.Open("todo.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var todoList TodoList

	json.Unmarshal(byteValue, &todoList)

	return todoList, err
}

func saveJson(todoList TodoList) {
	fmt.Println("Saving ToDo List to JSON")
	file, _ := json.MarshalIndent(todoList, "", " ")

	_ = ioutil.WriteFile("todo.json", file, 0644)

}

func printTodoList(todoList TodoList) {
	fmt.Println("Here is your Todo List:")
	for i := 0; i < len(todoList.TodoList); i++ {
		fmt.Print("Task # : " +
			strconv.Itoa(todoList.TodoList[i].TaskId) + " | ")
		fmt.Print("Description:  " + todoList.TodoList[i].Task + " | ")
		fmt.Println(" Status: " + todoList.TodoList[i].Status + " | ")
	}
}

func addTask(item string, todoList TodoList) TodoList {
	todoList.TodoList = append(
		todoList.TodoList,
		Todo{
			getMaxId(todoList),
			item,
			"incomplete"})
	return todoList
}

func deleteTask(id int, todoList TodoList) TodoList {
	for i, task := range todoList.TodoList {
		if task.TaskId == id {
			//todoList.TodoList = append(todoList.TodoList[:i], todoList.TodoList[i+1]...)
			fmt.Printf("Deleting task: ")
			fmt.Println(task)
			todoList.TodoList = append(todoList.TodoList[:i], todoList.TodoList[i+1:]...)
		}
	}
	//printTodoList(todoList)
	return todoList
}

func updateTaskStatus(id int, todoList TodoList) TodoList {

	//if no -status flag entered, ask for status ID listed above
	for i, task := range todoList.TodoList {
		if task.TaskId == id {
			fmt.Println(
				"Updating Task Status for " + strconv.Itoa(task.TaskId) + ": " + task.Task)
			fmt.Printf("update task status ['Incomplete', 'In Progress', 'Complete']: ")
			reader := bufio.NewReader(os.Stdin)
			text, _, _ := reader.ReadLine()
			todoList.TodoList[i].Status = string(text)

		}
	}
	return todoList
}

func getMaxId(todoList TodoList) int {
	//send back the highest value ID and increment
	var maxVal int = 0
	for _, task := range todoList.TodoList {
		if task.TaskId > maxVal {
			maxVal = task.TaskId
		}
	}
	return maxVal + 1
}

func main() {
	addItem := flag.String("add", "task description", "a string var")
	viewList := flag.Bool("list", false, "view task list")
	deleteItem := flag.Int("delete", 999, "task ID")
	updateItem := flag.Int("update", 999, "taskId")

	flag.Parse()

	var todoList TodoList

	var err error
	todoList, err = parseJSON()
	if err != nil {
		fmt.Println("Add your first task: ")
		reader := bufio.NewReader(os.Stdin)
		text, _, _ := reader.ReadLine()
		todoList = addTask(string(text), todoList)
		saveJson((todoList))
	}

	if *viewList {
		printTodoList(todoList)
	}

	if isFlagPassed("add") {
		todoList = addTask(*addItem, todoList)
		saveJson(todoList)
		printTodoList((todoList))
	}

	if isFlagPassed("delete") {
		todoList = deleteTask(*deleteItem, todoList)
		saveJson((todoList))
		printTodoList(todoList)
	}

	if isFlagPassed("update") {
		printTodoList(todoList)
		todoList = updateTaskStatus(*updateItem, todoList)
		saveJson(todoList)
		fmt.Println("Task Status for " + strconv.Itoa(*updateItem) + " updated")
	}
}
