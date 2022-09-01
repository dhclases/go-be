package main

import "fmt"

func main() {
	fmt.Println("hi")

	var palabra string = "esternocleidomastoideo"

	tampalabra := len(palabra)

	fmt.Println("Cantidad de letras =", tampalabra)

	for i := 0; i < tampalabra; i++ {
		fmt.Println(string(palabra[i]))

	}

}
