package models

import (
    "time"
	"github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// SearchRecord represents the user_search_record table structure
type SearchRecord struct {
    ID          uuid.UUID  // Using uuid.UUID for ID field
    SearchQuery string     `json:"search_query"`
    UserID      uuid.UUID  `json:"user_id"`
    SearchTime  time.Time  `json:"search_time"`
}

type SearchRecordHandlers interface {
	GetAllSearchRecords(c *gin.Context)
	GetSearchRecord(c *gin.Context)
	CreateSearchRecord(c *gin.Context)
}

type SearchRecordService interface {
	GetAllSearchRecords() ([]*SearchRecord, error)
	GetSearchRecord(uuid.UUID) (*SearchRecord, error)
	CreateSearchRecord(sr *SearchRecord) (uuid.UUID, error)
	DeleteAllSearchRecords() error
}

type SearchRecordRepository interface {
	GetAll() ([]*SearchRecord, error)
	Get(uuid.UUID) (*SearchRecord, error)
	Create(p *SearchRecord) (uuid.UUID, error)
	DeleteAll() error
}
