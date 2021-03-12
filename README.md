# GuíaNet Théa

Proyecto para servir las páginas para el catalogo de Théa.

## Contenido
- [Tecnologías](#tecnologías)
- [Herramientas](#herramientas)
- [Desarrollo](#desarrollo)
- [Deploy](#deploy)
- [Variables de entorno](#variables-de-entorno)

## Tecnologías
- Golang v 1.16
- Postgres

## Herramientas
- [golang-migrate](https://github.com/golang-migrate/migrate): Se usa para mantener el control de las versiones de la base de datos.

## Desarrollo.
- Se requiere instalar [Golang](https://golang.org/) y [Postgres](https://www.postgresql.org/), además de las herramientas especificadas.

- Se crea una base de datos con cualquier nombre y configurarlo, ya sea en una [variable de entorno](#variables-de-entorno) o directamente en el modulo _connection.go_ en la linea 15 sustituyendo la existente. 
```
instance, e := sql.Open("postgres", GetEnvVar("DB_URL", "postgres://postgres:abc123@localhost:5432/mexico_test_1?sslmode=disable"))
```

- Teniendo instalado [golang-migrate](https://github.com/golang-migrate/migrate) se ejecuta dentro de la carpeta que contiene el proyecto. Se sustituye [url] por la url de la base de datos
```
migrate -path db/migrations -database [url] -verbose up
```

- Se agregan los datos de prueba que se encuentran en db/data. [Esta entrada de StackOverflow](https://stackoverflow.com/questions/3204274/importing-sql-file-on-windows-to-postgresql) puede ayudar.

- Para correr el programa se necesita una terminal que se encuentre dentro de la carpeta que contiene el proyecto y ejecutar
```
go run main.go
```

## Deploy

- Se realiza un build del proyecto con el siguiente comando que generará un ejecutable. Este será dependiendo del Sistema Operativo que se este usando, asi que es necesario ejecutarlo en un sistema operativo como el del servidor.
```
go build
```

- Se requiere realizar la configuración de las [variables de entorno](#variables-de-entorno)

- Se levanta la base de datos de la misma forma que en desarrollo

- Se corre el ejecutable y debe mostrar el mensaje en consola.
```
Servidor iniciado
```

## Variables de entorno

- **DB_URL**: Indica la url de la base de datos en la que se guardará la información. Esta contempla el siguiente formato:
```
postgres://[user]:[password]@[host]:[port]/[database]?[...options]
```

- **PORT**: Indica el puerto donde estará servida la información. Se reciben numeros enteros y que sean puertos validos. Por ejemplo: ```8000```, ```8080```, ```9000```, ```3000```, etc

- **SECRET-KEY**: Indica la llave secreta con la que las sesiones serán creadas. 