package main

import "fmt"

func main() {

	var task1 = "Watch night movie"
	var task2 = "Read Book"
	var task3 = "Add Book"
	var task4 = "Write code"

	var taskIteams = []string{task1, task2, task3, task4}
	fmt.Println("##### Welcome TodoList App! #####\n")

	for index, task := range taskIteams {
		fmt.Println(index+1, task)
	}

}
