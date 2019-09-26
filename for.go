package main
import (
	"fmt"
	"strconv")

func main() {
	nombre, apellidoPaterno, apellidoMaterno, edad, colegiatura, aprobacion:= "MILTON JEONATHAN", "MEJIA", "FABIAN", 26, 3000, true
	meses:= [12]string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}

	if colegiatura > 2000 && aprobacion{
		fmt.Println("Nombre: " + nombre + " " + apellidoPaterno + " " + apellidoMaterno + "\nEdad: " + strconv.Itoa(edad))
		
		for i := 0; i < 12; i++ {
			colegiatura:= colegiatura + (colegiatura / 100 * 20)
			fmt.Println("Pago en " + meses[i] + ": " + strconv.Itoa(colegiatura))
		}

	} else {
		fmt.Println("Insuficiente")
	}
}