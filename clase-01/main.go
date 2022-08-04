// Este es un comentario simple

/*
   Este es un commetario multiple
*/

// Aqui definimos el punto de entrada principal: El paquete mai
package main

import (
	"fmt"
)

// Asi importamos las libreria. modo simple
/*
   import "fmt"

   el modo agrupado es
	import (
		"fmt"
		"http"
		"time"
	)

*/

// Definicion de structura: Esto lo vemos luego en detalle
type User struct {
	name       string
	occupation string
}

// Este proceso corren antes de commenzar el main()
/*
func init() {
	fmt.Println("Inicializando el programa principal")
}

*/

func main() {
	/*
		fmt.Println("Hola Digital")
		fmt.Printf("Esto es un numero:%v \n", 2323232)
		fmt.Printf("Esto es un numero:%T \n", 2323232)
	*/
	/*
	   Declaracion de variables
	*/

	// **** Variables enteras ***
	/*
		int, se asigna de acuerdo al SO (32 o 64 bits)
		int8,
		int16
		int32
		int64

	*/
	/*
		// Entero explicitamente
		var miEntero int = 1
		fmt.Printf("Mientero: %d \n", miEntero)
		fmt.Printf("Mientero es del tipo: %T \n", miEntero)

		// Entero8 explicitamente
		var miEntero8 int8 = 8
		fmt.Printf("Mientero8 : %d \n", miEntero8)
		fmt.Printf("Mientero8 es del tipo: %T \n", miEntero8)

		// Entero16 explicitamente
		var miEntero16 int16 = 16
		fmt.Printf("Mientero16: %d \n", miEntero16)
		fmt.Printf("Mientero16 es del tipo: %T \n", miEntero16)

		// Entero32 explicitamente
		var miEntero32 int32 = 32
		fmt.Printf("Mientero32: %d \n", miEntero32)
		fmt.Printf("Mientero32 es del tipo: %T \n", miEntero32)

		// Entero64 implicitament
		miEntero7 := 64
		fmt.Printf("Mientero7: %d \n", miEntero7)
		fmt.Printf("Mientero7 es del tipo: %T \n", miEntero7)

	*/
	/*
			// **** Variables float ***

			var miFloat32 float32 = 32.02
			fmt.Printf("Mi float32: %v \n", miFloat32)
			fmt.Printf("Mi flaot32 es del tipo: %T \n", miFloat32)

			var miFloat64 float64 = 345452.02545
			fmt.Printf("Mi float64: %v \n", miFloat64)
			fmt.Printf("Mi flaot64 es del tipo: %T \n", miFloat64)

			// **** Variable boolean ***

			estaActivo := true
			fmt.Printf("Mi estaActivo: %v \n", estaActivo)
			fmt.Printf("Mi estaActivo es del tipo: %T \n", estaActivo)

		// *** Definir una constante
		const tipo_usuario = "admin"
		fmt.Printf("Mi tipo_usuario: %v \n", tipo_usuario)
		fmt.Printf("Mi tipo_usuario es del tipo: %T \n", tipo_usuario)

	*/
	/*
		// ** Declaracion en bloque
		var (
			esCliente bool = false
			monto          = 0.0
		)
		fmt.Printf("Mi esCliente: %v  y monto: %v \n", esCliente, monto)

		// ** Declaracion en bloque
		const (
			constante1 = false
			constante2 = 0.0
			constante3 = "esta es un string"
		)
		fmt.Printf("constante1: %v , constante2: %v constante3: %v \n", constante1, constante2, constante3)
				// ** Declaracion de string
	*/
	/*
		var cadena1 string = "esta es la cadena 1"
		cadena2 := "esta es la cadena 2"
		fmt.Printf("cadena1: %v , cadena2: %v \n", cadena1, cadena2)

		// ** Declaracion multiples
		var a, b, c = 5.25, 25.25, 14.15 // Multiple float32
		fmt.Println(a, b, c)

	*/
	/*
		// *** Arreglos / Array
		var intArreglos [5]int

		vals := []int{1, 2, 3, 4, 5}

		var y [5]int = [5]int{10, 20, 30} // Asignacion parcial

		x := [...]int{10, 20, 30} // Inicializacion con ellipses ...

		z := [5]int{1: 10, 3: 30} // Inicializar con indice especifico

		// Asignacion
		intArreglos[0] = 100

		// Acceso
		fmt.Printf("Valor de intArreglos[0]: %v \n", intArreglos[0])
		fmt.Printf("Valor de y : %v \n", y[0])
		fmt.Printf("Valor de x : %v \n", x[1])
		fmt.Printf("Valor de z : %v \n", z[3])
		fmt.Printf("Valores: %v \n", vals[0])
		fmt.Printf("Tipo de dato de intArreglos : %T \n", intArreglos)

		fmt.Printf("Valores: %v \n", y[3])
	*/

	/*
		// Declaracin implicita

		// *** Estructura / Struct
		u := User{"John Doe", "gardener"}
		fmt.Printf("Usuario: %v \n", u.name)
	*/

	// *** Slices
	/*
		n Slice es una "ventana" a un Array que se compone de 3 cosas.

		(1) Puntero: el puntero se utiliza para apuntar al primer elemento del Array al que se puede acceder a través del slice.
		(2) Longitud: La longitud es el número total de elementos presentes en el Array.
		(3) Capacidad: La capacidad representa el tamaño máximo del Array.

	*/
	/*
		var intSliceVacio []int // Inicializar vacio
		fmt.Printf("intSlice \tLen: %v \tCap: %v \tTipo: %T \n", len(intSliceVacio), cap(intSliceVacio), intSliceVacio)

		// Declaracion con Make
		var intSlice02 = make([]int, 10)        // La misma longitud y capacidad
		var strSlice02 = make([]string, 10, 20) // Con diferente longitud y capacidad

		fmt.Printf("intSlice02 \tLen: %v \tCap: %v \tTipo: %T \n", len(intSlice02), cap(intSlice02), intSlice02)
		fmt.Printf("strSlice02 \tLen: %v \tCap: %v \tTipo: %T \n", len(strSlice02), cap(strSlice02), strSlice02)

		// Actualizar un item
		intSlice02[0] = 555

		// Agregar un item
		intSlice02 = append(intSlice02, 369)

		fmt.Printf("intSlice02 \tLen: %v \tCap: %v \tTipo: %T \n", len(intSlice02), cap(intSlice02), intSlice02)
		fmt.Printf("intSlice02 \titem 0 : %v \n", intSlice02[0])
		fmt.Printf("intSlice02 \titem 0 : %v \n", intSlice02[0])
	*/
	/*
		var numSlice = []int{10, 20, 30, 40}
		var paisesSlice = []string{"India", "Canada", "Japan"}

		fmt.Printf("numSlice \tLen: %v \tCap: %v\n", len(numSlice), cap(numSlice))
		fmt.Printf("paisesSlice \tLen: %v \tCap: %v\n", len(paisesSlice), cap(paisesSlice))
		fmt.Printf("paisesSlice \titem 0 : %v \n", paisesSlice[0])
		fmt.Printf("paisesSlice \titem 1 : %v \n", paisesSlice[1])

		var intSlice = []int{10, 20, 30, 40}
		var strSlice = []string{"India", "Canada", "Japan"}

		fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
		fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	*/
	/*
		// *** Mapas/Map similares a Java HashTable Python Dic
		//                  Clave Valor
		var employee = map[string]int{"Mark": 10, "Daniela": 20} // Definicion explicita
		fmt.Printf("employee \tItems: %v \n", employee)
		fmt.Printf("employee \tLen: %v \tValor de Daniela: %v\n", len(employee), employee["Daniela"])
		// Con la funcion make

		var employee2 = make(map[string]int)
		employee2["Julio"] = 10
		employee2["Pedro"] = 20
		employee2["Sabrina"] = 30

		fmt.Printf("employee2 \tLen: %v \tvalor de Pedro: %v\n", len(employee2), employee2["Pedro"])

		paises := map[string]string{
			"sk": "Slovakia",
			"ru": "Rusia",
			"de": "Alemania",
			"no": "Noruega",
		} // Defincion Implicita
		fmt.Printf("Paises: %v \n", paises["de"])
	*/

	// *** Conversion de Datos
	/*
		// *** Int to String
		miEntero2 := 300
		varString1 := strconv.Itoa(miEntero2)
		fmt.Printf("varString1: %v \n", varString1)
		fmt.Printf("miEntero: %v \n", miEntero2)

		// *** String to int

		strVar := "100"
		// NOTA: Recepcion multiple de valores: El valor convertido y un error si ocurriese
		intVar, err := strconv.Atoi(strVar)


		fmt.Printf("varString2: %v \n", strVar)
		fmt.Printf("entero: %d. error: %v \n", intVar, err)

		strVar2 := "2"

		intVar8, err := strconv.ParseInt(strVar2, 0, 8)
		fmt.Printf(" intVar8 : %v,   error: %v, Tipo: %T \n", intVar8, err, intVar)

		intVar16, err := strconv.ParseInt(strVar, 0, 16)
		fmt.Printf(" intVar16: %v, error: %v, Tipo: %T \n", intVar16, err, intVar)

		intVar32, err := strconv.ParseInt(strVar, 0, 32)
		fmt.Printf(" intVar32: %v, error: %v, Tipo: %T \n", intVar32, err, intVar)

		intVar64, err := strconv.ParseInt(strVar, 0, 64)
		fmt.Printf(" intVar64: %v, error: %v, Tipo: %T \n", intVar64, err, intVar)
	*/

	/*
				Los operadores de go son similares al resto de los lenguajes.

		+, suma
		-, resta
		*, multiplicación
		/, división
		<, menor que
		<=, menor o igual que
		>, mayor que
		>=, mayor o igual que
		%, el módulo o residuo
		!=, inequivalencia
		==, igualdad
		!, negación
		&&, operador AND
		||, operador OR
		++, incremental
		--, decremental


	*/

	// *** Condicionales
	/*
		if  condition {
		    // Aqui colocamos el codigo
		}

	*/
	/*
		// *** estilos de condicionales

		x := true
		if x {
			fmt.Println("La condicion es cierta")
		}

		numero := 200
		if numero == 200 {
			fmt.Println("Numero es igual a " + strconv.Itoa(numero))
		} else {
			fmt.Println("Numero no es igual a" + strconv.Itoa(numero))
		}

		if numero == 50 {
			fmt.Println("Es igual a 50 ")
		} else if numero == 100 {
			fmt.Println("Es igual a 100")
		} else {
			fmt.Println("El numero es:" + strconv.Itoa(numero))
		}
	*/
	/*
		today := time.Now()

		switch today.Day() {
		case 1:
			fmt.Println("Hoy es primero. ")
		case 5:
			fmt.Println("Hoy es el quinto. ")
		case 10:
			fmt.Println("Hoy es 10. ")
		case 15:
			fmt.Println("Hoy es 15. ")
		case 25:
			fmt.Println("Hoy es e25. ")
		case 31:
			fmt.Println("Hoy es 31. ")
		default:
			fmt.Println("Sin informacion.")
		}

		i := 45

		switch {
		case i < 10:
			fmt.Println("i es menos de 10")
		case i < 50:
			fmt.Println("i es menos de 50")
		case i < 100:
			fmt.Println("i es menos de 100")
		}

		fmt.Println(" Segunda evaluacion ")

		i2 := 45
		switch {
		case i2 < 10:
			fmt.Println("i es menos de 10")
			fallthrough // Evalua la siguiente condicion
		case i2 < 50:
			fmt.Println("i es menos de 50")
			fallthrough
		case i2 < 100:
			fmt.Println("i es menos de 100")
		}
	*/
	// *** Bucles / Loop
	const max = 3

	for k := 1; k <= max; k++ {
		fmt.Printf("Bucle 1 k=%v \n", k)
	}
	fmt.Println(" *** Fin Bucle 1")

	k := 1

	for ; k <= max; k++ {
		fmt.Printf("Bucle 2 k=%v \n", k)
	}
	fmt.Println(" *** Fin Bucle 2")

	k = 1
	for k <= max {
		fmt.Printf("Bucle 3 k=%v \n", k)
		k++
	}
	fmt.Println(" *** Fin Bucle 3")

	for k := 1; ; k++ {
		fmt.Printf("Bucle 4 k=%v \n", k)
		if k == max {
			break
		}
	}
	fmt.Println(" *** Fin Bucle 4")

}
