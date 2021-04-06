package models

import "time"

type LoginRegistryData struct {
	InitDate  time.Time `json:"initDate" validate:"required"`
	FinalDate time.Time `json:"finalDate" validate:"required"`
}

func (l LoginRegistryData) GetValidableData() interface{} {
	return &l
}
