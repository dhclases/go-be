package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) ([]string, error) {
	res, err := os.ReadFile(path)
	data := strings.Split(string(res), ";")
	return data, err

}

func main() {

	var (
		total         float64
		totalProducto float64
		tab           string = "\t"
	)

	data, err := readFile("./data.csv")

	if err != nil {
		fmt.Println("Error en la lectua del archivo")
	}

	for i := 0; i < len(data)-1; i++ {

		var line = strings.Split(data[i], ",")

		// Encabezado o tiems de la linea
		for j := 0; j < len(line); j++ {
			fmt.Printf("%s\t\t", line[j])
		}

		// Etiqueta de total
		if i == 0 {
			fmt.Printf("%s", "Total")
		}

		// Calcular los subtotales y total
		if i != 0 {

			precio, err := strconv.ParseFloat(line[1], 64)
			if err != nil {
				fmt.Println("No se pudo parsear el precio")
			}

			cantidad, err := strconv.ParseFloat(line[2], 64)
			if err != nil {
				fmt.Println("No se pudo parsear la cantidad")
			}

			totalProducto = precio * cantidad
			total += totalProducto

		}

		// Imprimir el subtotal
		if i != 0 {
			fmt.Printf("\t%v", totalProducto)
		}

		fmt.Print("\n")

	}

	// Imprimir el total
	fmt.Printf(strings.Repeat("*", 60) + "\n")
	fmt.Printf(strings.Repeat(tab, 6)+"Total: %.2f\n", total)

}

// Luego
// * Separar por funciones y responsabilidades
// * Formatear las cifras eje 20.00
