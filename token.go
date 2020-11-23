package base

import (
	"context"
	"github.com/VaulGuard/base/dto"
)

type TokenStorage interface {
	Get(context.Context, interface{}) (dto.TokenDto, error)
	Create(context.Context, *dto.TokenDto) (*dto.TokenDto, error)
}

type TokenService interface {
	Generate(context.Context, interface{}) string
	Verify(context.Context, string) (dto.ApplicationDto, bool)
}
