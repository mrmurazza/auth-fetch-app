package user

import (
	"time"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

type User struct {
	ID          int
	Phonenumber string
	Name        string
	Password    string
	Role        Role
	CreatedAt   *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   *time.Time `gorm:"default:current_timestamp"`
	DeletedAt   *time.Time `gorm:"default: null"`
}
