package main

import (
	"actividad/clase/handlers"
	"actividad/clase/store"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// Definir el router de chequeo
	handlers.NewHealthRouter(engine)

	// Inicializa producto
	productList := store.LoadProducts("products.json")

	handlers.NewProductRouter(engine, productList)

	// Iniciar el servidor por el puerto 8080
	engine.Run(":8080")

}
