package main

import (
	"fmt"
	"math"
)

// Go Interface - `Shape`
type Shape interface {
	Area() float64
	Perimeter() float64
}

/*
Rectangle
*/
type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

/*
   Circle
*/

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

// Range >> iterate over key, value
func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	return totalArea
}

type UserDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (ud UserDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range ud.shapes {
		totalArea += s.Area()
	}
	return totalArea
}

func main() {

	// Polimorfismo
	printHeader("*** Interfaces *** ")

	var s Shape = Circle{5.0}
	fmt.Printf("Shape Tipo = %T, Shape Valor = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	s = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Tipo = %T, Shape Valor = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)

	printHeader("*** Interfaces como Tipo *** ")

	drawing := UserDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{3, 5},
			Rectangle{4, 7},
		},
		bgColor: "red",
		fgColor: "white",
	}

	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())

	// Polimorfismo
	printHeader("*** Polimorfismo *** ")

	var s2 Shape

	s2 = Circle{5}
	fmt.Printf("(%v, %T)\n", s2, s2)
	fmt.Printf("Shape area = %v\n", s.Area())

	s2 = Rectangle{4, 7}
	fmt.Printf("(%v, %T)\n", s2, s2)
	fmt.Printf("Shape area = %v\n", s.Area())

}

func printHeader(s string) {
	fmt.Println()
	fmt.Println(s)
	fmt.Println()

}
