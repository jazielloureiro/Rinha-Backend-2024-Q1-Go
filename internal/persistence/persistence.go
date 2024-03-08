package persistence

import "github.com/jazielloureiro/Rinha-Backend-2024-Q1-Go/internal/entity"

type AccountRepository interface {
	Get(int) (entity.Account, error)
	Update(*entity.Account, int) error
}
