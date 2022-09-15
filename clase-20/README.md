# consignas-go-db

## Despliegue

Para empezar la ejecucion de la apliocac, ejecutamos el siguiente comando ubicados en la raiz del proyecto

<pre><code> go run cmd/server/main.go </code></pre>

Commandos Util

docker run -d -p 33060:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=secret mysql