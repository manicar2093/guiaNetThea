package models

import "time"

type LoginRegistryData struct {
	InitDate  time.Time `json:"initDate" validate:"required"`
	FinalDate time.Time `json:"finalDate" validate:"required, gtefield=InitDate"`
}

func (l LoginRegistryData) GetValidableData() interface{} {
	return &l
}
