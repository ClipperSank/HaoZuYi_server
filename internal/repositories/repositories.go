package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
)

type Repositories struct {
	Users      		models.UserRepository
	SearchRecords  models.SearchRecordRepository
	Posts      		models.PostRepository
	Messages   		models.MessageRepository
	Locations 		models.LocationRepository
	Categories 		models.CategoryRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Users:      	NewUserRepository(db),
		SearchRecords: NewSearchRecordRepository(db),
		Posts:      	NewPostRepository(db),
		Messages:   	NewMessageRepository(db),
		Locations:  	NewLocationRepository(db),
		Categories: 	NewCategoryRepository(db),
	}
}
