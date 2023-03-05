package methods

import (
	"acme/internal/users/domain"

	"gorm.io/gorm"
)

type repository struct {
	gormDriver *gorm.DB
}

// NewRepository returns a new instance of the transaction ledger repository.
func NewRepository(gormDriver *gorm.DB) domain.Repository {
	return &repository{
		gormDriver: gormDriver,
	}
}
