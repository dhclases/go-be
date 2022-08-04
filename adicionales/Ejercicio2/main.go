package main

import "fmt"

func main() {
	var Mes int = 1

	Meses := make(map[int]string)

	MesesNombres := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	for i := 0; i < len(MesesNombres); i++ {
		Meses[i] = MesesNombres[i]
	}

	fmt.Printf("El # mes: %v y su nombre: %v \n", Mes, Meses[Mes-1])

}
