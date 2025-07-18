package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("##### Welcome TodoList App! #####\n")

	http.HandleFunc("/", helloUser)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {

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
