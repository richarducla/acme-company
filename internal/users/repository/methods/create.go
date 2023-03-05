package methods

import (
	"acme/internal/models"
	rmodels "acme/internal/users/repository/models"

	"github.com/labstack/gommon/log"
)

// Create adds a new record to the user table.
func (r *repository) Create(userModel models.User) (models.User, error) {
	dao := rmodels.ToUserDAO(userModel)

	err := r.gormDriver.Model(dao).Create(&dao).Error
	if err != nil {
		log.Error("error creating user %w: ", err)

		return models.User{}, err
	}

	user, err := r.Get(dao.Username)
	if err != nil {
		log.Error("error retrieving user %w: ", err)

		return models.User{}, err
	}

	return user, nil
}
