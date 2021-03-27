package utils

import (
	"fmt"
	"os"
)

// GetEnvVar obtiene una variable de entorno o regresa el string especificado
func GetEnvVar(env string, possible string) string {
	d, ok := os.LookupEnv(env)
	if !ok {
		return possible
	}
	return d
}

// GetPortFromEnvVar valida la existencia del puerto en el enviroment y lo regresa con el formato necesario
func GetPortFromEnvVar(envVar, possible string) string {
	p := GetEnvVar(envVar, possible)
	if p != possible {

		return fmt.Sprintf(":%s", p)
	}
	return fmt.Sprintf(":%s", possible)
}
