package dto

import "time"

type TokenDto struct {
	ID            interface{}
	Value         []byte
	ApplicationId interface{}
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Application   ApplicationDto
}
