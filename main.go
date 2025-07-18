package main

import (
	"fmt"
	"net/http"
)

var task1 = "Watch night movie"
var task2 = "Read Book"
var task3 = "Add Book"
var task4 = "Write code"
var taskIteams = []string{task1, task2, task3, task4}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTask)

	http.ListenAndServe(":8080", nil)
}

func showTask(writer http.ResponseWriter, request *http.Request) {
	for index, task := range taskIteams {
		fmt.Fprintln(writer, index+1, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to our TodoList App!"
	fmt.Fprintln(writer, greeting)
}

func printTask(taskIteams []string) {
	fmt.Println("List of my task")
	for index, task := range taskIteams {
		fmt.Printf("%d: %s\n\n", index+1, task)
	}
}

func addTask(taskIteams []string, newTask string) []string {
	var updatedTaskIteams = append(taskIteams, newTask)
	return updatedTaskIteams
}
