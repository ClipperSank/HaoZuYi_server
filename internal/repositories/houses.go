package repositories

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"fyno/server/internal/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type houseRepository struct {
	DB *sql.DB
}

func NewHouseRepository(db *sql.DB) models.HouseRepository {
	return &houseRepository{
		DB: db,
	}
}

func (hr *houseRepository) GetAll() ([]*models.House, error) {
	query := `SELECT h.id, h.user_id, h.kind, h.name, h.age, h.gender, h.content, l.id AS location_id, l.name AS location_name, c.id AS category_id, c.name AS category_name, h.created_at
				FROM houses AS h
				JOIN locations AS l ON h.location_id = l.id 
				JOIN categories As c ON h.category_id = c.id 
				ORDER BY h.created_at DESC`

	fmt.Println("query", query)
	rows, err := hr.DB.Query(query)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer rows.Close()

	var houses []*models.House
	for rows.Next() {
		var h models.House
		err := rows.Scan(&h.ID, &h.UserID, &h.Kind, &h.Name, &h.Age, &h.Gender, &h.Content, &h.Location.ID, &h.Location.Name, &h.Category.ID, &h.Category.Name, &h.CreatedAt)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		houses = append(houses, &h)
	}
	fmt.Println("houses", houses)
	return houses, nil
}

func (hr *houseRepository) Get(id uuid.UUID) (*models.House, error) {
	fmt.Println("id", id)
	query := `SELECT h.id, h.user_id, h.kind, h.name, h.age, h.gender, h.content, l.id AS location_id, l.name AS location_name, c.id AS category_id, c.name AS category_name, h.created_at, ARRAY_AGG((pi.url, pi.rank)) AS house_images
				FROM houses AS h
				JOIN locations AS l ON h.location_id = l.id 
				JOIN categories As c ON h.category_id = c.id
				LEFT JOIN house_images AS pi ON h.id = pi.house_id
				WHERE h.id = $1
				GROUP BY h.id, l.id, c.id`

	row := hr.DB.QueryRow(query, id)

	var h models.House
	var houseImages []HouseImages
	err := row.Scan(&h.ID, &h.UserID, &h.Kind, &h.Name, &h.Age, &h.Gender, &h.Content, &h.Location.ID, &h.Location.Name, &h.Category.ID, &h.Category.Name, &h.CreatedAt, pq.Array(&houseImages))
	if err != nil {
		fmt.Println("error", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("houseImages", houseImages)
	fmt.Println("h", h)
	h.HouseImages = make([]models.HouseImages, len(houseImages))
	for i, pi := range houseImages {
		h.HouseImages[i] = models.HouseImages(pi)
	}
	return &h, nil
}

func (hr *houseRepository) Create(h *models.House) (uuid.UUID, error) {
	query := `INSERT INTO houses (id, user_id, kind, name, age, gender, content, location_id, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := hr.DB.Exec(query, h.ID, h.UserID, h.Kind, h.Name, h.Age, h.Gender, h.Content, h.Location.ID, h.Category.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return h.ID, nil
}

func (hr *houseRepository) DeleteAll() error {
	query := `DELETE FROM houses`
	_, err := hr.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}