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

## Endpoints

Este sistema cuenta con varios endpoints que a continuación se describen:

**WEB**

* _/_ _GET_
Este endpoint renderiza el template on_dev.html

* _/index_ _GET_
Renderiza el template login.html

* _/{page}_ _GET_
Dependiendo de la página que se solicite renderiza el html que corresponda

* _/login_ _POST_
Endpoint para realizar el login al sistema

* _/logout/_ _GET_
Realiza la salida de sistema

**ADMIN**

**NOTA: Todos comienzan con _/admin_**

* _/ GET_
Renderiza el index del administrador

* _/user/all GET_
Renderiza el template users.html

* _/user/registry GET_
Renderiza el template registry_user.html

* _/logginRegistry GET_
Renderiza el template bitacora.html

* _/user/{idUser} GET_
Renderiza el template edit_user.html

* _/user/registry POST_
Realiza el registro de un usuario. Detalles del JSON para el request:
```
RolID            int32  `json:"rol_id" validate:"required"`
Name             string `json:"name" validate:"required"`
PaternalSureName string `json:"paternal_surename" validate:"required"`
MaternalSureName string `json:"maternal_surename"`
Email            string `json:"email" validate:"required,email"`
Password         string `json:"password" validate:"required"`
PasswordConfirm  string `json:"password_confirm" validate:"required,eqfield=Password"`
```

* _/user/delete/{idUser} DELETE_
Pone en status false al usuario que se solicite

* _/user/restore_password PUT_
Realiza el cambio de contraseña de un usuario. Detalles del JSON para el request:
```
ID              int32  `json:"id" validate:"required"`
Password        string `json:"password" validate:"required"`
PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
```

* _/user/update PUT_
Realiza la actualización de un usuario, menos su contraseña. Detalles del JSON para el request:
```
ID               int32  `json:"id" validate:"required"`
RolID            int32  `json:"rol_id" validate:"required"`
Name             string `json:"name" validate:"required"`
PaternalSureName string `json:"paternal_surename" validate:"required"`
MaternalSureName string `json:"maternal_surename"`
Email            string `json:"email" validate:"required,email"`
```

* _/catalogs/{catalog} GET_
Obtiene los datos de los catalogos registrados. Todos contienen la siguiente información:
```
ID          int    `json:"id"`
Description string `json:"description"`
```

* _/login_registry/create POST_
Crea un .xlsx con la información del logueo de los usuarios a la plataforma. Detalles del JSON para el request:
```
InitDate  time.Time `json:"initDate" validate:"required"`
FinalDate time.Time `json:"finalDate" validate:"required"`
```

## Variables de entorno

- **DB_URL**: Indica la url de la base de datos en la que se guardará la información. Esta contempla el siguiente formato:
```
postgres://[user]:[password]@[host]:[port]/[database]?[...options]
```

- **PORT**: Indica el puerto donde estará servida la información. Se reciben numeros enteros y que sean puertos validos. Por ejemplo: ```8000```, ```8080```, ```9000```, ```3000```, etc

- **SECRET-KEY**: Indica la llave secreta con la que las sesiones serán creadas. 