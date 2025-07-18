package main

import "fmt"

var task1 = "Watch night movie"
var task2 = "Read Book"
var task3 = "Add Book"
var task4 = "Write code"

var taskIteams = []string{task1, task2, task3, task4}

func main() {
	fmt.Println("##### Welcome TodoList App! #####\n")
	printTask()
}

func printTask() {
	fmt.Println("List of my task")
	for index, task := range taskIteams {
		fmt.Printf("%d: %s\n\n", index+1, task)
	}
}
