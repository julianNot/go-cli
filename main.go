package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/julianNot/go-cli/tasks"
)

func instructions() {
	fmt.Println("Use: go-clid-crud [ list|add|complete|delete ]")
}

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		instructions()
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What is your Task?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.Add(tasks, name)
		task.SaveTask(file, tasks)
		fmt.Println(tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Enter [id] to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be a number")
		}
		tasks = task.Delete(tasks, id)
		task.SaveTask(file, tasks)
	}

}
