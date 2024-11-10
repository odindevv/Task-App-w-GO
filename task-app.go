package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var TASKS = make(map[int]string)
var scanner = bufio.NewScanner(os.Stdin)
var idCounter = 1

func main() {
	for {
		clearConsole()
		fmt.Println("Bienvenidos al sistema de Tareas!")
		showMenu()
		option := getOption()

		switch option {
		case 1:
			addTask()
		case 2:
			editTask()
		case 3:
			deleteTask()
		case 4:
			viewTasks()
		case 5:
			fmt.Println("Salida del programa con exito!")
			return
		default:
			fmt.Println("Opcion no valida, intenta de nuevo.")
		}
		pauseConsole()
	}

}

func showMenu() {
	fmt.Println("¿Que desea realizar? (Elija una opcion)")
	fmt.Println("1) Crear tareas")
	fmt.Println("2) Editar tarea")
	fmt.Println("3) Eliminar tarea")
	fmt.Println("4) Ver tareas")
	fmt.Println("5) Salir")
}

func getOption() int {
	scanner.Scan()
	option, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return -1
	}
	return option
}

func addTask() {
	clearConsole()
	fmt.Println("Escriba el nombre de la tarea (o 'q' para regresar al menu):")
	scanner.Scan()
	nameTask := scanner.Text()
	if nameTask == "q" {
		return
	}
	TASKS[idCounter] = nameTask
	fmt.Printf("Tarea '%s', agregada con ID: %d\nEnter para continuar.\n", nameTask, idCounter)
	idCounter++
}

func editTask() {
	clearConsole()
	if len(TASKS) == 0 {
		fmt.Println("No hay tareas para editar.")
		return
	}

	viewTasks()
	fmt.Println("Ingrese el ID de la tarea a editar (o 'q' para regresar al menu):")
	idTask := getTaskID()
	if idTask == -1 {
		return
	}

	fmt.Printf("Tarea actual: %s\nEscriba el nuevo nombre de la tarea  (o 'q' para regresar al menu):\n", TASKS[idTask])
	scanner.Scan()
	newName := scanner.Text()
	if newName == "q" {
		return
	}
	TASKS[idTask] = newName
	fmt.Println("Tarea actualizada correctamente")
}

func deleteTask() {
	clearConsole()
	if len(TASKS) == 0 {
		fmt.Println("No hay tareas para eliminar.")
		return
	}

	viewTasks()
	fmt.Println("Ingrese el ID de la tarea a elminar  (o 'q' para regresar al menu):")
	idTask := getTaskID()
	if idTask == -1 {
		fmt.Println("ID invalido o regreso al menu.")
		return
	}

	delete(TASKS, idTask)
	fmt.Println("Tarea eliminada correctamente.")
}

func viewTasks() {
	clearConsole()
	if len(TASKS) == 0 {
		fmt.Println("No hay tareas registradas.")
	} else {
		fmt.Println("Tareas registradas:")
		for id, task := range TASKS {
			fmt.Printf("%d) %s\n", id, task)
		}
	}
}

func getTaskID() int {
	scanner.Scan()
	input := scanner.Text()
	if input == "q" {
		return -1
	}

	idTask, err := strconv.Atoi(input)
	if err != nil || TASKS[idTask] == "" {
		return -1
	}
	return idTask
}

func clearConsole() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pauseConsole() {
	fmt.Println("Presione enter para continuar...")
	scanner.Scan()
}
