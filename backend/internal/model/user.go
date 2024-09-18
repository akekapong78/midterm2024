package model

import "github.com/akekapong78/workflow/internal/constant"

type User struct {
	ID       uint              `json:"id" gorm:"primaryKey"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Role     constant.UserRole `json:"role"`
}

type RequestLogin struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
}

type RequestUser struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Role     constant.UserRole `json:"role"`
}

type ResponseUser struct {
	ID       uint              `json:"id"`
	Username string            `json:"username"`
	Role     constant.UserRole `json:"role"`
}
