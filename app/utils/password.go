package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordUtils interface {
	HashPassword(password []byte) (string, error)
	ValidatePassword(hashed, password string) error
}

type PasswordUtilsImpl struct{}

func NewPasswordUtils() *PasswordUtilsImpl {
	return &PasswordUtilsImpl{}
}

// HashPassword crea un hash con el algoritmo seleccionado para el sistema
func (p PasswordUtilsImpl) HashPassword(password []byte) (string, error) {
	b, e := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if e != nil {
		return "", fmt.Errorf("An error occurred creating user's password")
	}
	return string(b), nil
}

// ValidatePassword valida que la contraseña recibida coincida con el hash proporcionado
func (p PasswordUtilsImpl) ValidatePassword(hashed, password string) error {
	e := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if e != nil {
		return fmt.Errorf("Invalid password")
	}
	return nil
}
