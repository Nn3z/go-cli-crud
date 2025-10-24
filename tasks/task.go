package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Task representa una tarea simple con ID, título y estado de completado.
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ListTasks imprime la lista de tareas en consola.
// Muestra un indicador ✓ si la tarea está completada.
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
		// Imprime ID, título y estado
		fmt.Printf("%d %s [%s]\n", task.ID, task.Title, status)
	}
}

// AddTask crea una nueva tarea y la agrega al slice existente.
// Devuelve el slice actualizado.
func AddTask(tasks []Task, name string) []Task {

	newTask := Task{
		ID:        ObtenerSiguienteID(tasks), // asigna el siguiente ID disponible
		Title:     name,
		Completed: false,
	}
	return append(tasks, newTask)

}

// SaveTasksToFile serializa las tareas a JSON y sobreescribe el archivo dado.
// Usa bufio.Writer para un escritura eficiente y asegura que el archivo quede truncado antes de escribir.
func SaveTasksToFile(tasks []Task, file *os.File) {
	// Convierte el slice de tareas a JSON
	bytes, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	// Mueve el cursor al inicio del archivo para sobreescribir
	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	// Trunca el archivo para eliminar contenido previo que pudiera quedar
	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	// Escribe los bytes JSON usando un buffer
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)

	if err != nil {
		panic(err)
	}

	// flush asegura que todo se escribe físicamente en el archivo
	err = writer.Flush()

	if err != nil {
		panic(err)
	}

}

// DeleteTask elimina la tarea con el ID especificado y devuelve el slice actualizado.
// Si no se encuentra el ID, devuelve el slice original.
func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			// Construye un nuevo slice excluyendo el elemento en i
			return append(tasks[:i], tasks[i+1:]...)
		}
	}

	return tasks
}

// CompleteTask marca la tarea con el ID dado como completada.
// Devuelve el slice actualizado.
func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}
	return tasks
}

// ObtenerSiguienteID calcula el próximo ID a usar.
// Si no hay tareas devuelve 1; en caso contrario usa el último ID + 1.
func ObtenerSiguienteID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}

	return tasks[len(tasks)-1].ID + 1
}
