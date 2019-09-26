package main
import (
	"fmt"
	"strconv")

func main() {
	nombre, apellidoPaterno, apellidoMaterno, edad, colegiatura:= "MILTON JEONATHAN", "MEJIA", "FABIAN", 26, 3000
	fmt.Println("Nombre: " + nombre + " " + apellidoPaterno + " " + apellidoMaterno + " " + "\nEdad: " + strconv.Itoa(edad) +
	"\nColegiatura: " + strconv.Itoa(colegiatura))
}