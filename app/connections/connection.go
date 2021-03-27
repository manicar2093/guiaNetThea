package connections

import (
	"database/sql"
	"log"

	// Driver de la base de datos
	_ "github.com/lib/pq"
	"github.com/manicar2093/guianetThea/app/utils"
)

// DB es la instancia para la conexión a la base de datos
var DB *sql.DB

func init() {
	instance, e := sql.Open("postgres", utils.GetEnvVar("DB_URL", "postgres://postgres:abc123@localhost:5432/mexico_test_1?sslmode=disable"))
	if e != nil {
		log.Fatal("Error al conectar la base de datos. Detalles: \n", e)
	}

	e = instance.Ping()
	if e != nil {
		log.Fatal("No hay respuesta en el ping. Detalles: \n", e)
	}

	DB = instance
}
