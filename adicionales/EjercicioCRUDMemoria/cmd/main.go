package main

import (
	"actividad/clase/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	engine.SetTrustedProxies([]string{"127.0.0.1"})

	// Definir el router
	handler.NewHealthRouter(engine)
	handler.NewProductRouter(engine, *handler.InitProductHandler())
	handler.NewProductRouter(engine, *handler.InitProductHandler())

	// Iniciar el servidor por el puerto 8080
	engine.Run(":9001")

}
