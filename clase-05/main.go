package main

import (
	"fmt"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func estaVacia(m Matrix) bool {

	if len(m.valores) == 0 {
		fmt.Println("Matrix vacía")
		return true
	}
	return false
}

func (m Matrix) Set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones especificadas.")
	}

}

func (m Matrix) Cuadratica() bool {
	if (m.alto == m.ancho) && m.alto != 0 {
		return true
	}
	return false
}

func (m Matrix) Max() float64 {

	if estaVacia(m) {
		fmt.Println("Matrix vacía")
		return 0.0
	}

	max := m.valores[0]

	for _, elemento := range m.valores {
		if elemento > max {
			max = elemento
		}
	}
	return max
}

// {1, 2, 3, 4, 54, 65, 76, 87, 87}
//  - ancho-

func (m Matrix) Print() {
	if estaVacia(m) {
		return
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}
}
func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 54, 65, 76, 87, 87},
		alto:    3,
		ancho:   3,
	}

	// Forma 1
	fmt.Println("Forma 1")
	m.Set()
	fmt.Println("Visualizar matriz")
	m.Print()
	fmt.Printf("Es cuadrática: %v \n", Matrix.Cuadratica(m))
	m.Cuadratica()
	fmt.Printf("Cual es el máximo: %v \n", Matrix.Max(m))
	m.Cuadratica()

	// Forma 2
	fmt.Println("***")
	fmt.Println("Forma 2")
	Matrix.Set(m)
	fmt.Println("Visualizar matriz")
	Matrix.Print(m)
	fmt.Printf("Es cuadrática: %v \n", Matrix.Cuadratica(m))
	Matrix.Cuadratica(m)
	fmt.Printf("Cual es el máximo: %v \n", Matrix.Max(m))
	Matrix.Cuadratica(m)

}
