package base

import (
	"context"
	"github.com/VaulGuard/base/dto"
	"sync"
)

type SecretService interface {
	Paginate(ctx context.Context, applicationID interface{}, page, perPage int) (map[string]string, error)
	Get(ctx context.Context, applicationID interface{}, key []string) (map[string]string, error)
	GetOne(ctx context.Context, applicationID interface{}, key string) (dto.SecretDto, error)
	Create(ctx context.Context, applicationID interface{}, key, value string) (dto.SecretDto, error)
	Update(ctx context.Context, applicationID interface{}, key, newKey, value string) (dto.SecretDto, error)
	Delete(ctx context.Context, applicationID interface{}, key string) error
	InvalidateCache(ctx context.Context, applicationID interface{}) error
}

type Service struct {
	mutex             *sync.RWMutex
	cacheLimit        int
	cache             [1024]map[string]dto.SecretDto
	encryptionService Encryption
}
