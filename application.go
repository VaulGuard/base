package base

import (
	"context"
	"github.com/VaulGuard/base/dto"
)

type ApplicationService interface {
	List(context.Context, int, func([]dto.ApplicationDto) error) error
	GetByName(context.Context, string) (dto.ApplicationDto, error)
	Create(context.Context, string) (dto.ApplicationDto, error)
	Get(context.Context, int, int) ([]dto.ApplicationDto, error)
	GetOne(context.Context, interface{}) (dto.ApplicationDto, error)
	Update(context.Context, interface{}, string) (dto.ApplicationDto, error)
	Delete(context.Context, interface{}) error
}
