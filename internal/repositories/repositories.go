package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
)

type Repositories struct {
	Users      		models.UserRepository
	SearchRecords  models.SearchRecordRepository
	// Houses			models.HouseRepository
	// Contracts      models.ContractRepository
	// Reviews			models.ReviewRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:      	NewUserRepository(db),
		SearchRecords: NewSearchRecordRepository(db),
		// Houses:		   NewHouseRepository(db),
		// Contracts:	   NewContractsRepository(db),
		// Reviews: 		NewReviewRepository(db),
	}
}
