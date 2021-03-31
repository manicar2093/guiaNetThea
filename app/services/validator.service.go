package services

import "github.com/manicar2093/guianetThea/app/models"

type ValidatorService interface {
	Validate(data models.Validable) ([]models.ErrorValidationDetail, bool)
}

type ValidatorServiceImpl struct {
}

func NewValidatorService() ValidatorService {
	return &ValidatorServiceImpl{}
}

func (v ValidatorServiceImpl) Validate(data models.Validable) ([]models.ErrorValidationDetail, bool) {
	panic("not implemented") // TODO: Implement
}
