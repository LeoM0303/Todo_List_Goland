package main

import "fmt"

func main() {
	fmt.Println("##### Welcome TodoList App! #####\n")

	var task1 = "Watch night movie"
	var task2 = "Read Book"
	var task3 = "Add Book"
	var task4 = "Write code"

	var taskIteams = []string{task1, task2, task3, task4}

	printTask(taskIteams)
	fmt.Println()

	//place for add task
	taskIteams = addTask(taskIteams, "Add taks")
	taskIteams = addTask(taskIteams, "Practice coding in Go")

	fmt.Println("___Updated list___")
	fmt.Println()
	printTask(taskIteams)
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
