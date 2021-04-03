package services

import (
	"testing"

	"github.com/manicar2093/guianetThea/app/models"
	"github.com/stretchr/testify/assert"
)

// TestValidate valida el happy path del metodo
func TestValidate(t *testing.T) {
	toValidate := models.CreateUserData{
		RolID:            1,
		Name:             "Name",
		PaternalSureName: "Test",
		MaternalSureName: "Test",
		Email:            "test@test.com",
		Password:         "holi",
		PasswordConfirm:  "holi",
	}

	validator := NewValidatorService()
	errors, valid := validator.Validate(toValidate)

	assert.True(t, valid, "La estructura debió ser valida")
	assert.Len(t, errors, 0, "No debió regresar ningun error")
}

func TestValidateWithError(t *testing.T) {
	toValidate := models.CreateUserData{
		RolID:            1,
		Name:             "",
		PaternalSureName: "Test",
		MaternalSureName: "Test",
		Email:            "testtest.com",
		Password:         "holi",
		PasswordConfirm:  "hol",
	}

	validator := NewValidatorService()
	errors, valid := validator.Validate(toValidate)

	assert.False(t, valid, "La estructura no debió ser valida")
	assert.Len(t, errors, 3, "No debió regresar ningun error")
}
