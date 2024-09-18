package user

import (
	"github.com/akekapong78/workflow/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}

func (r Repository) GetUsers() ([]model.ResponseUser, error) {
	var users []model.ResponseUser
	err := r.Database.Table("users").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r Repository) GetUser(id string) (model.ResponseUser, error) {
	var user model.ResponseUser
	err := r.Database.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return model.ResponseUser{}, err
	}

	return user, nil
}

func (r Repository) CreateUser(req *model.RequestUser) (uint, error) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
		Role: req.Role,
	}
	err := r.Database.Table("users").Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r Repository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.Database.Table("users").Where("username = ?", username).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}