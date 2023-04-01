package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	task "github.com/julianNot/go-cli/tasks"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)
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
	}	else {
		tasks = []task.Task{}
	}

	fmt.Println(tasks)

}
