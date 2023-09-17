package models

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID  `json:"id"`
	IndexPage      string     `json:"indexpage"`
	Username       string     `json:"username"`
	Role           string     `json:"role"`
	CreatedAt      time.Time  `json:"created_at"`
	Account        string     `json:"account"`
	Password       string     `json:"password"`
	Age            int        `json:"age"`
	Birthday       time.Time  `json:"birthday"`
	ContractCount  int        `json:"contract_count"`
	HousesForRent  int        `json:"houses_for_rent"`
	OwnedHouses    int        `json:"owned_houses"`
}

type UserHandlers interface {
	GetUser(c *gin.Context)
	GetUserByName(c *gin.Context)
	CreateUser(c *gin.Context)
	// TODO: UpdateUser(c * gin.Context)
}

type UserService interface {
	GetUser(uuid.UUID) (*User, error)
	GetUserByName(string) (*User, error)
	CreateUser(u *User) (*User, error)
	// TODO: UpdateUser(uuid.UUID, u *User) (*User, error)
	DeleteAllUsers() error
}

type UserRepository interface {
	Get(uuid.UUID) (*User, error)
	GetByName(string) (*User, error)
	Create(u *User) (*User, error)
	// TODO: Update(uuid.UUID, u *User) (*User, error)
	DeleteAll() error
}
