package services

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/manicar2093/guianetThea/app/models"
)

var validate *validator.Validate

type ValidatorService interface {
	Validate(data models.Validable) ([]models.ErrorValidationDetail, bool)
}

type ValidatorServiceImpl struct {
}

func NewValidatorService() ValidatorService {
	return &ValidatorServiceImpl{}
}

func (v ValidatorServiceImpl) Validate(data models.Validable) ([]models.ErrorValidationDetail, bool) {
	var e []models.ErrorValidationDetail
	err := validate.Struct(data.GetValidableData())
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var m models.ErrorValidationDetail
			m.Field = err.Field()
			switch err.ActualTag() {
			case "required":
				m.Constraints = append(m.Constraints, "is required")
				fallthrough
			case "eqfield":
				m.Constraints = append(m.Constraints, fmt.Sprintf("must be equals to %s", err.Param()))
			case "email":
				m.Constraints = append(m.Constraints, "must have email format")
			default:
				m.Constraints = append(m.Constraints, "somethig happend. please contact support")
			}
			e = append(e, m)
		}

		return e, false
	}

	return e, true

}

func init() {
	validate = validator.New()
}
