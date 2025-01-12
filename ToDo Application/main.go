package main

import (
	"fmt"
	"net/http"
)

var taskOne = "Watch Go Crash Course"
var myFirstApp = "Make first application"
var dayEnd = "Done for day" //short version of var declaration
var taskItems = []string{taskOne, myFirstApp, dayEnd}

func main() {
	//############### Code without HTTP Requests ##################
	//var taskOne = "Watch Go Crash Course"
	//var myFirstApp = "Make first application"
	//var dayEnd = "Done for day" //short version of var declaration
	//var taskItems = []string{taskOne, myFirstApp, dayEnd}
	//fmt.Println("###### Welcome to our Todo Application! ######")
	//printTasks(taskItems)
	//fmt.Println()
	//taskItems = addTasks(taskItems, "Go for a run")
	//taskItems = addTasks(taskItems, "Practice coding in Go")
	//fmt.Println()
	//printTasks(taskItems)
	////fmt.Println("Tasks:", taskItems[0], taskItems[1], taskItems[2])

	//############### Code with HTTP Requests ##################
	//fmt.Println("###### Welcome to our Todo Application! ######")
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		//fmt.Println(index+1, ":", task)
		fmt.Fprintln(writer, task)
		//fmt.Fprintf("%d: %s\n", index+1, task)
		//
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello World! Welcome to our ToDo List Application!"
	fmt.Fprintf(writer, greeting)
}

func printTasks(taskItems []string) { //passing a task item slice to the function
	fmt.Println("List of my ToDo's!")
	for index, task := range taskItems {
		//fmt.Println(index+1, ":", task)
		fmt.Printf("%d: %s\n", index+1, task)

	}
}

func addTasks(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	printTasks(updatedTaskItems)
	return updatedTaskItems
}
