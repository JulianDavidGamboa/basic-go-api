package model

import "errors"

var (
	// ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil     = errors.New("La persona no puede ser nula")
	ErrIDPersonDoesNotExists = errors.New("El ID de la persona no existe")
)
