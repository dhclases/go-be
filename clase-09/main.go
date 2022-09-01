package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer recoverPanic()
	defer callCerrarProceso()

	// callError()

	// callEnv()
	/*
		_, err := callCustomError()
		if err != nil {
			fmt.Println(err)
		}
	*/
	// callPanic()

	fmt.Println("*******")
}

func callError() {
	error_1 := fmt.Errorf(" Error type 1")
	error_2 := fmt.Errorf(" Error type 2 : %w", error_1)

	err := errors.Unwrap(error_2)
	if &err == &error_1 {
		fmt.Println("exactly the same error")
	} else {
		fmt.Println("same error")
	}

	fmt.Printf("unwrapped error: %v\n", err)
}

func callEnv() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	jh, ok := os.LookupEnv("JAVA_HOME")
	if !ok {
		fmt.Println("JAVA_HOME no esta definida", jh)
	}
}

type CustomError struct{}

func (ce *CustomError) Error() string {
	return "Custom Error bla bla"
}

func callCustomError() (string, error) {
	return "", &CustomError{}
}

func callPanic() {
	panic("Un super error")
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando de un Panic. Error:\n", r)
	} else {
		fmt.Println("Todo bien no hubo panico")
	}
}

func callCerrarProceso() {
	fmt.Println("Cerrando el proceso")
}
