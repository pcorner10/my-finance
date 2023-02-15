package controller

import "github.com/pcorner10/my-finance/api/models"

func GetUserByUserName(user *models.User) (*models.User, error) {
	user, err := user.GetUserByUsername()
	if err != nil {
		return nil, err
	}

	return user, nil
}
