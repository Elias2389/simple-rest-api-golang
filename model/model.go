package model

import "errors"

var (
	// ErrPersonCanNotBeNil
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// ErrIDPersonDoesNotExists
	ErrIDPersonDoesNotExists = errors.New("la persona no existe")
)
