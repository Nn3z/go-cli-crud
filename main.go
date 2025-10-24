package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/Nn3z/go-cli-crud/tasks"
)

func main() {
	// Abre (o crea si no existe) el archivo tasks.json con permisos de lectura/escritura.
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		// Panic en caso de error al abrir/crear el archivo.
		panic(err)
	}

	// Asegura que el archivo se cierre al finalizar la ejecución.
	defer file.Close()

	var tasks []task.Task

	// Obtiene información del archivo para saber si está vacío.
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// Si el archivo no está vacío, leer y deserializar JSON en la variable tasks.
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		// Si está vacío, iniciar slice vacío para evitar nil.
		tasks = []task.Task{}
	}

	// Validación básica de argumentos de línea de comandos.
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Discriminador principal de comandos: list, add, delete, complete...
	switch os.Args[1] {
	case "list":
		// Muestra la lista de tareas por consola.
		task.ListTasks(tasks)
	case "add":
		// Lee desde stdin el nombre de la nueva tarea.
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es tu tarea?")
		name, _ := reader.ReadString('\n')

		name = strings.TrimSpace(name)

		// Añade la tarea y guarda los cambios al archivo.
		tasks = task.AddTask(tasks, name)
		task.SaveTasksToFile(tasks, file)

	case "delete":
		// Verifica que se proporcione el ID a eliminar.
		if len(os.Args) < 3 {
			fmt.Println("Por favor proporciona el ID de la tarea a eliminar.")
			return
		}

		// Convierte el argumento a entero y valida.
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalido. Por favor proporciona un numero valido.")
			return
		}

		// Elimina la tarea y guarda los cambios.
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasksToFile(tasks, file)
	case "complete":
		// Similar a delete: valida el ID y marca la tarea como completada.
		if len(os.Args) < 3 {
			fmt.Println("Por favor proporciona el ID de la tarea a eliminar.")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID invalido. Por favor proporciona un numero valido.")
			return
		}

		// Marca la tarea como completada y guarda los cambios.
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasksToFile(tasks, file)
	default:
		// Si el comando no coincide, mostrar uso.
		printUsage()
	}
}

// printUsage muestra las opciones disponibles del CLI.
func printUsage() {
	fmt.Println("TIENES ESTAS OPCIONES PARA USAR ESTE CLI: go-cli-crud [list | add | update | delete ]")
}
