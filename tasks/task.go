package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {

		status := " "

		if task.Completed == true {
			status = "✓"
		}
		fmt.Printf("%d %s [%s]\n", task.ID, task.Title, status)
	}
}

func AddTask(tasks []Task, name string) []Task {

	newTask := Task{
		ID:        ObtenerSiguienteID(tasks),
		Title:     name,
		Completed: false,
	}
	return append(tasks, newTask)

}

func SaveTasksToFile(tasks []Task, file *os.File) {
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

	// flush asegura que todo se escribe en el archivo
	err = writer.Flush()

	if err != nil {
		panic(err)
	}

}

func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}

	return tasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}
	return tasks
}

func ObtenerSiguienteID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
