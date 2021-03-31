package models

type ErrorValidationDetail struct {
	Field       string   `json:"field"`
	Constraints []string `json:"constraints"`
}

type Validable interface {
	// GetValidableData permite compartir la estructura que se debe validar. Este struct debe llevar los tags correspondientes del paquete go-playground/validator.
	GetValidableData() interface{}
}
