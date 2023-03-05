package domain

import (
	"acme/internal/models"
)

//go:generate mockery --name=Repository --filename=user.go --structname=UserRepository --output=../../../test/mocks/repositories --outpkg=repositories

// Repository Base methods for transaction ledger repository
type Repository interface {
	Create(user models.User) (models.User, error)
	Get(username string) (models.User, error)
}
