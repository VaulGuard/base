package dto

import "time"

type ApplicationDto struct {
	ID        interface{}
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
