package main

import (
	"errors"
	"fmt"
)

const (
	min  = "minimo"
	max  = "maximo"
	prom = "promedio"
	sum  = "suma"
	rest = "resta"
)

func main() {
	fmt.Println("Inicio")

	// Calcular min

	min1 := calcularMin(2, 5, 6, 7, 8, 9, 2)
	fmt.Println("Valor min1", min1)

	// Calcular max
	max1 := calcularMax(2, 5, 6, 7, 8, 9, 2)
	fmt.Println("Valor max2", max1)

	// Calcular prom

	prom1 := calcularProm(2, 5, 6, 7, 8, 9, 2)
	fmt.Println("Valor prom1", prom1)

	// Definir funciones

	// Func min
	minFunc, err := funcionFactory(min)

	if err != nil {
		fmt.Println(err)

	}

	// Func max
	maxFunc, err := funcionFactory(max)

	if err != nil {
		fmt.Println(err)

	}

	// Func prom
	promFunc, err := funcionFactory(prom)

	if err != nil {
		fmt.Println(err)

	}

	valorMin := minFunc(2, 5, 6, 7, 8, 9, 2)
	fmt.Printf("Calculo min: (1) %v, (2) %v \n", min1, valorMin)

	valorMax := maxFunc(2, 5, 6, 7, 8, 9, 2)
	fmt.Printf("Calculo max: (1) %v, (2) %v \n", max1, valorMax)

	valorProm := promFunc(2, 5, 6, 7, 8, 9, 2)
	fmt.Printf("Calculo prom: (1) %v, (2) %v \n", prom1, valorProm)

	oper := funcionFactory2("div")
	r := oper(2, 2)
	fmt.Printf("Calculo r: %v", r)

}

func calcularMin(nums ...int) int {

	if len(nums) == 0 {
		return 0
	}

	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}

	return min

}

func calcularMax(nums ...int) int {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max

}

func calcularProm(nums ...int) int {

	if len(nums) == 0 {
		return 0
	}

	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]

	}

	return total / len(nums)

}

func funcionFactory(operacion string) (func(...int) int, error) {

	switch operacion {
	case min:
		return calcularMin, nil
	case max:
		return calcularMax, nil
	case prom:
		return calcularProm, nil
	default:
		return nil, errors.New("no existe la operación")
	}
}

func opSuma(v1, v2 float64) float64 {
	return v1 + v2
}

func opResta(v1, v2 float64) float64 {
	return v1 + v2
}

func funcionFactory2(operacion string) func(v1, v2 float64) float64 {

	switch operacion {
	case min:
		return opSuma
	case max:
		return opResta

	default:
		return nil
	}
}
