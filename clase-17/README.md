# Agregar la siguiente validacion

## Usar variable dentorno

TOKEN=clave-re-segura

``` golang

token := c.GetHeader("TOKEN")

if token == "" {
    web.Failure(c, 401, errors.New("token not found"))
    return
}

if token != os.Getenv("TOKEN") {
    web.Failure(c, 401, errors.New("invalid token"))
    return
}
```

Utilizar el componente para el manejo de errores personalizados en response
