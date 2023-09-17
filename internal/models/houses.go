package models

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type House struct {
    ID            uuid.UUID `json:"id"`
    UserID        uuid.UUID `json:"user_id"`
    Address       string    `json:"address"`
    IsRenting     int       `json:"is_renting"`
    Price         float64   `json:"price"`
    Size          int       `json:"size"`
    Kitchen       int       `json:"kitchen"`
    Bathroom      int       `json:"bathroom"`
    SleepingRoom  int       `json:"sleeping_room"`
    CreatedAt     time.Time `json:"created_at"`
}

type HouseHandlers interface {
	GetAllHouses(c *gin.Context)
	GetHouse(c *gin.Context)
	CreateHouse(c *gin.Context)
    // TODO: UpdateHouse(c *gin.Context)
}

type HouseService interface {
	GetAllHouses() ([]*House, error)
	GetHouse(uuid.UUID) (*House, error)
	CreateHouse(sr *House) (uuid.UUID, error)
	// TODO: UpdateHouse(sr *House) (uuid.UUID, error)
	DeleteAllHouses() error
}

type HouseRepository interface {
	GetAll() ([]*House, error)
	Get(uuid.UUID) (*House, error)
	Create(p *House) (uuid.UUID, error)
	// TODO: Update(p *House) (uuid.UUID, error)
	DeleteAll() error
}
