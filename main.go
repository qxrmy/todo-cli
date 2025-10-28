package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func clearConsole() {
	fmt.Print("\033[2J\033[1;1H")
}

func main() {
	_ = LoadTasks()

	reader := bufio.NewReader(os.Stdin)
	var choice int
	var description string
	var taskID int
	var saveTask string

	for {
		color.Cyan("=== TODO CLI ===")
		color.White("1. add new task | 2. tasks list | 3. do task")
		color.White("4. delete task | 5. edit task | 6. exit")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}
		switch choice {
		case 1:
			fmt.Println("describe the task:")
			description, err = reader.ReadString('\n')
			description = strings.TrimSpace(description)
			if description == "" {
				color.Yellow("description is empty")
				continue
			}
			newTask := Task{Description: description, Done: false}
			(&newTask).Add()
		case 2:
			TasksList()
		case 3:
			TasksList()
			fmt.Println("enter task id: ")
			_, err := fmt.Scanln(&taskID)
			if err != nil {
				color.Red("invalid input")
				continue
			}
			i := FindIndexByID(taskID)
			if i == -1 {
				color.Red("task not found")
				continue
			}
			(&tasks[i]).DoTask()
		case 4:
			TasksList()
			fmt.Println("enter task id: ")
			_, err := fmt.Scanln(&taskID)
			if err != nil {
				color.Yellow("invalid input")
				continue
			}
			DeleteTask(taskID)
		case 5:
			TasksList()
			fmt.Println("enter task id: ")
			_, err := fmt.Scanln(&taskID)
			if err != nil {
				color.Yellow("invalid input")
				continue
			}
			fmt.Println("enter new description: ")
			description, err = reader.ReadString('\n')
			description = strings.TrimSpace(description)
			if description == "" {
				color.Yellow("description is empty")
				continue
			}
			EditTask(taskID, description)
		case 6:
			fmt.Println("save tasks?(y/n): ")
			_, err = fmt.Scanln(&saveTask)
			if saveTask == "y" {
				if err := SaveTasks(); err != nil {
					color.Red("save error:", err)
					return
				}
				color.Green("tasks saved")
				return
			}
			fmt.Println("tasks are not saved")
			fmt.Println("exit...")
			return
		}
		fmt.Println("Press Enter to continue...")
		_, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		clearConsole()
	}
}
