package tasks

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No Tasks")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf("[%s] %d %s \n", status, task.ID, task.Name)
	}
}

func Add(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNexId(tasks),
		Name:     name,
		Complete: false,
	}

	return append(tasks, newTask)
}

func Delete(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func GetNexId(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
