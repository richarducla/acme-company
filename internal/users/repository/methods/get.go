package methods

import (
	"acme/internal/models"
	rmodels "acme/internal/users/repository/models"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

// Get find an user by username.
func (r *repository) Get(username string) (models.User, error) {

	dao := rmodels.User{}
	err := r.gormDriver.Where("username = ?", username).First(&dao).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("error user not found %w: ", err)

			return models.User{}, err
		}

		return models.User{}, err
	}

	return rmodels.ToUserModel(dao), nil
}
