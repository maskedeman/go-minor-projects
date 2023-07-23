package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct{
	ID int
	Name string
	Complete bool
	Delete bool
}
var tasks []Task

func add_task(name string) {
taskID := len(tasks)+1
task:= Task{
	ID: taskID,
	Name: name,
	Complete: false,

}
tasks=append(tasks,task)
fmt.Printf("Task '%s' is added with ID %d\n",name,taskID)
}

func listTasks(){
	if len(tasks)==0{
		fmt.Println("No tasks found")
		return
	}
	fmt.Println("Task List:")
	for _,task := range tasks{
		status:="Incomplete"
		if task.Complete{
			status="Completed"
		}
		fmt.Printf("[%d]%s-%s\n",task.ID,task.Name,status)
}}
func completeTask(taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Complete = true
			fmt.Printf("Task '%s' marked as completed.\n", task.Name)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID)
}
func deleteAll(){
	tasks=nil
fmt.Println("All Tasks deleted")}
func deleteTask(taskID int){
	for i, task:= range tasks{
		if task.ID == taskID{
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task '%s' deleted.\n", task.Name)
			return
		}
	}
		
		fmt.Printf("No such id exists in list of tasks: %d.",taskID)
	}
	


	

func main(){
	fmt.Println("To-Do List App")
	fmt.Println("Enter 'add <task_name>' to add a new task.")
	fmt.Println("Enter 'list' to view all tasks.")
	fmt.Println("Enter 'complete <task_id>' to mark a task as completed.")
	fmt.Println("Enter 'delAll' to delete all tasks.")
	fmt.Println("Enter 'del <task_id>' to delete a task.")
	fmt.Println("Enter 'exit' to quit the app.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		command := strings.Fields(input)
		if len(command) < 1 {
			continue
		}

		switch command[0] {
		case "add":
			if len(command) > 1 {
				taskName := strings.Join(command[1:], " ")
				add_task(taskName)
			} else {
				fmt.Println("Usage: add <task_name>")
			}
		case "list":
			listTasks()
		case "complete":
			if len(command) == 2 {
				taskID := command[1]
				var id int
				_, err := fmt.Sscanf(taskID, "%d", &id)
				if err != nil {
					fmt.Println("Invalid task ID. Please provide a valid integer.")
					continue
				}
				completeTask(id)
			} else {
				fmt.Println("Usage: complete <task_id>")
			}
		case "delAll":
			deleteAll()
		case "del":
			if len(command)==2{
				taskID:=command[1]
				var id int
				_,err:=fmt.Sscanf(taskID,"%d",&id)
				
				if err !=nil{
					fmt.Println("Invalid Task Id")
					
					continue
				}
				deleteTask(id)
			}else{
fmt.Println("Usage: del <task_id>")
			}
			
		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
	fmt.Println("Goodbye!")

}