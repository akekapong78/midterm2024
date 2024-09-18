package utils

import (
	"github.com/akekapong78/workflow/internal/model"
	"gorm.io/gorm"
)

func GetUserIdByUsername(username string, db *gorm.DB) (uint, error) {
	var user model.User
	err := db.Table("users").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}