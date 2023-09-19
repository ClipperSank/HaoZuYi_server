package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Contract struct {
	ID            uuid.UUID `json:"id"`
	RenterID      uuid.UUID `json:"renter_id"`
	LandlordID    uuid.UUID `json:"landlord_id"`
	HouseID       uuid.UUID `json:"house_id"`
	ContractText  string    `json:"contract"`
	Rent          float64   `json:"rent"`
	StartTime     time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time,omitempty"`
	RenterReview  *string   `json:"renter_review,omitempty"`
	LandlordReview *string   `json:"landlord_review,omitempty"`
}

type ContractHandlers interface {
	GetAllContracts(c *gin.Context)
	GetContract(c *gin.Context)
	CreateContract(c *gin.Context)
}

type ContractService interface {
	GetAllContracts() ([]*Contract, error)
	GetContract(uuid.UUID) (*Contract, error)
	CreateContract(p *Contract) (uuid.UUID, error)
	DeleteAllContracts() error
}

type ContractRepository interface {
	GetAll() ([]*Contract, error)
	Get(uuid.UUID) (*Contract, error)
	Create(p *Contract) (uuid.UUID, error)
	DeleteAll() error
}
