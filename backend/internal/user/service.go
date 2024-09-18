package user

import (
	"errors"

	"github.com/akekapong78/workflow/internal/auth"
	"github.com/akekapong78/workflow/internal/model"
	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	secret     string
}

func NewService(db *gorm.DB, secret string) Service {
	return Service{
		Repository: NewRepository(db),
		secret:     secret,
	}
}

func (s Service) GetUsers() ([]model.ResponseUser, error) {
	users, err := s.Repository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s Service) GetUser(id string) (model.ResponseUser, error) {
	user, err := s.Repository.GetUser(id)
	if err != nil {
		return model.ResponseUser{}, err
	}

	return user, nil
}

func (s Service) GetUserByUsername(username string) (model.ResponseUser, error) {
	user, err := s.Repository.GetUserByUsername(username)
	if err != nil {
		return model.ResponseUser{}, err
	}

	return model.ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Role: user.Role,
	}, nil
}

func (s Service) CreateUser(req *model.RequestUser) (model.ResponseUser, error) {
	// Hash password
	hash, err := auth.HashPassword(req.Password)
	if err != nil {
		return model.ResponseUser{}, err
	}

	req.Password = hash

	// Create user
	userId, err := s.Repository.CreateUser(req)
	if err != nil {
		return model.ResponseUser{}, err
	}

	return model.ResponseUser{
		ID: userId,
		Username: req.Username,
		Role: req.Role,
	}, nil
}


func (s Service) Login(req *model.RequestLogin) (string, error) {
	// Get user
	user, err := s.Repository.GetUserByUsername(req.Username)
	if err != nil {
		return "",  errors.New("invalid user or password")
	}

	// Check password
	if ok := auth.CheckPasswordHash(req.Password, user.Password); !ok {
		return "", errors.New("invalid password")
	}

	// Generate token
	token, err := auth.GenerateToken(user.Username, user.Role, s.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}