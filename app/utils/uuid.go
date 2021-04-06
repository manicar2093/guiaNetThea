package utils

import uuid "github.com/satori/go.uuid"

type UUIDGeneratorUtils interface {
	// CreateUUIDV4 crea un UUID V4 con el package uuid
	CreateUUIDV4() string
}

type UUIDGeneratorUtilsImpl struct{}

func NewUUIDGeneratorUtils() UUIDGeneratorUtils {
	return UUIDGeneratorUtilsImpl{}
}
func (u UUIDGeneratorUtilsImpl) CreateUUIDV4() string {
	return uuid.NewV4().String()
}
