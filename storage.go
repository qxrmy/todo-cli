package main

import (
	"encoding/json"
	"os"
)

func SaveTasks() error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func LoadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}
	if len(tasks) > 0 {
		maxID := 0
		for _, t := range tasks {
			if t.TaskID > maxID {
				maxID = t.TaskID
			}
		}
		nextID = maxID + 1
	}
	return nil
}
