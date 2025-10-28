package main

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
)

type Task struct {
	TaskID      int
	Description string
	Done        bool
}

var tasks []Task
var nextID = 1

func (task *Task) Add() {
	task.TaskID = nextID
	nextID++
	tasks = append(tasks, *task)
}

func TasksList() {
	if len(tasks) == 0 {
		color.Red("no tasks found")
		return
	}
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Done != tasks[j].Done {
			return !tasks[i].Done
		}
		return tasks[i].TaskID < tasks[j].TaskID
	})
	for i, task := range tasks {
		idColor := color.New(color.FgCyan).SprintFunc()
		descColor := color.New(color.FgYellow).SprintFunc()
		statusColor := color.New(color.FgGreen).SprintFunc()

		if !task.Done {
			descColor = color.New(color.FgHiYellow).SprintFunc()
			statusColor = color.New(color.FgRed).SprintFunc()
		}

		fmt.Printf("%d. %s %s %s\n", i+1, idColor(task.TaskID), descColor(task.Description),
			statusColor("(done: "+fmt.Sprint(task.Done)+")"))
	}
}

func (task *Task) DoTask() {
	task.Done = true
}

func DeleteTask(taskID int) {
	i := FindIndexByID(taskID)
	if i == -1 {
		color.Red("task not found")
		return
	}
	tasks = append(tasks[:i], tasks[i+1:]...)
}

func EditTask(taskID int, newDescription string) {
	i := FindIndexByID(taskID)
	if i == -1 {
		color.Red("task not found")
		return
	}
	tasks[i].Description = newDescription
}

func FindIndexByID(taskID int) int {
	if len(tasks) == 0 {
		return -1
	}
	for i := range tasks {
		if tasks[i].TaskID == taskID {
			return i
		}
	}
	return -1
}
