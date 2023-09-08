package models

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
}

type UserHandlers interface {
	GetUser(c *gin.Context)
	GetUserByName(c *gin.Context)
	CreateUser(c *gin.Context)
}

type UserService interface {
	GetUser(uuid.UUID) (*User, error)
	GetUserByName(string) (*User, error)
	CreateUser(u *User) (*User, error)
	DeleteAllUsers() error
}

type UserRepository interface {
	Get(uuid.UUID) (*User, error)
	GetByName(string) (*User, error)
	Create(u *User) (*User, error)
	DeleteAll() error
}
