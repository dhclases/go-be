package main

/*
Para instalar Gin ejecutar el siguiente commando:
go get -u github.com/gin-gonic/gin

Luego lo importamos  a nuestro código:
import "github.com/gin-gonic/gin"
*/

// go mod tidy

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type persona struct {
	Nombre    string
	Apellido  string
	Edad      int
	Direccion string
	Telefono  int
	Activo    bool
}

func printPersona(p persona) {
	h := "*"
	log.Print(strings.Repeat(h, 60))
	log.Print(p)
	log.Print(strings.Repeat(h, 60))

}
func main() {

	router := gin.Default()

	//Definir una persona de ejemplo
	jsonData := `{"Nombre": "Juan", "Apellido": "Perez", "Edad": 25, "Direccion": "Calle falsa 123","Telefono": 47852201, "Activo": true}`

	var p persona
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal("Error Fatal")
		log.Fatal(err)
	}

	//Imprimo la persona por consola
	printPersona(p)

	//Captura la solicitud GET “/persona"
	router.GET("/persona", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"persona": p,
		})
	})

	router.Run(":8080")

	//Ejecutar con: go run main.go
	//Para para el servidor hacer: Ctrl + C
}
