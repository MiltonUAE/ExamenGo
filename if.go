package main
import (
	"fmt"
	"strconv")

func main() {
	nombre, apellidoPaterno, apellidoMaterno, edad, colegiatura, aprobacion:= "MILTON JEONATHAN", "MEJIA", "FABIAN", 26, 3000, true

	if colegiatura > 2000 && aprobacion{
		fmt.Println(nombre + " " + apellidoPaterno + " " + apellidoMaterno + " " + strconv.Itoa(edad))
	} else {
		fmt.Println("Insuficiente")
	}
}