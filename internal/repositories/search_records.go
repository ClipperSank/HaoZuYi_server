package repositories

import (
	"database/sql"
	"fmt"
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type searchrecordRepository struct {
	DB *sql.DB
}

func NewSearchRecordRepository(db *sql.DB) models.SearchRecordRepository {
	return &searchrecordRepository{
		DB: db,
	}
}

func (srr *searchrecordRepository) GetAll() ([]*models.SearchRecord, error) {
	query := `SELECT id, search_query, user_id, search_time FROM user_search_record ORDER BY search_time DESC`

	fmt.Println("query", query)
	rows, err := srr.DB.Query(query)
	if err != nil {
		 fmt.Println("error", err)
		 return nil, err
	}
	defer rows.Close()

	var searchRecords []*models.SearchRecord
	for rows.Next() {
		 var sr models.SearchRecord
		 err := rows.Scan(&sr.ID, &sr.SearchQuery, &sr.UserID, &sr.SearchTime)
		 if err != nil {
			  fmt.Println("error", err)
			  return nil, err
		 }
		 searchRecords = append(searchRecords, &sr)
	}
	fmt.Println("searchRecords", searchRecords)
	return searchRecords, nil
}

func (srr *searchrecordRepository) Get(id uuid.UUID) (*models.SearchRecord, error) {
	fmt.Println("id", id)
	query := `
		 SELECT sr.id, sr.search_query, sr.user_id, sr.search_time
		 FROM user_search_record AS sr
		 WHERE sr.id = $1
	`

	row := srr.DB.QueryRow(query, id)

	var sr models.SearchRecord
	err := row.Scan(&sr.ID, &sr.SearchQuery, &sr.UserID, &sr.SearchTime)
	if err != nil {
		 fmt.Println("error", err)
		 if err == sql.ErrNoRows {
			  return nil, nil
		 }
		 return nil, err
	}

	return &sr, nil
}

func (srr *searchrecordRepository) Create(sr *models.SearchRecord) (uuid.UUID, error) {
	// Define the query to insert a new search record into the database
	query := `
		INSERT INTO user_search_record (id, search_query, user_id, search_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute the query and retrieve the ID of the newly inserted record
	var id uuid.UUID
	err := srr.DB.QueryRow(query, sr.ID, sr.SearchQuery, sr.UserID, sr.SearchTime).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (srr *searchrecordRepository) DeleteAll() error {
	query := `DELETE FROM user_search_record`
	_, err := srr.DB.Exec(query)
	if err != nil {
		 return err
	}

	return nil
}
