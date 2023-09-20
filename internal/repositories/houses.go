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

func (hr *houseRepository) GetAll(userID uuid.UUID) ([]*models.House, error) {
	query := `SELECT h.id, h.user_id, h.address, h.is_renting, h.price, h.size, h.kitchen, h.bathroom, h.sleeping_room, h.house_images, h.created_at
				FROM houses AS h
				WHERE h.user_id = $1
				ORDER BY h.created_at DESC`

	fmt.Println("query", query)
	rows, err := hr.DB.Query(query, userID)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer rows.Close()

	var houses []*models.House
	for rows.Next() {
		var h models.House
		err := rows.Scan(&h.ID, &h.UserID, &h.Address, &h.IsRenting, &h.Price, &h.Size, &h.Kitchen, &h.Bathroom, &h.SleepingRoom, &h.CreatedAt)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		houses = append(houses, &h)
	}
	fmt.Println("houses", houses)
	return houses, nil
}

type HouseImages models.HouseImage

func (hi *HouseImages) Scan(src interface{}) error {
	fmt.Println("Scanning")
	if src == nil {
		return errors.New("HouseImages Scan: null value")
	}

	// The src value returned by pq is of type []byte, so we need to convert it to a string
	srcStr, ok := src.([]byte)
	if !ok {
		return errors.New("HouseImages Scan: invalid type")
	}
	fmt.Println("srcStr", srcStr)
	srcStr = bytes.Trim(srcStr, "{}")
	fmt.Println("srcStr", srcStr)
	// Split the string into URL and rank using the separator ","
	parts := strings.Split(string(srcStr), ",")
	if len(parts) != 2 {
		return errors.New("HouseImages Scan: invalid format")
	}

	// Parse the URL and rank values from the string parts
	url := strings.Trim(parts[0], "\"()")
	fmt.Println("url", url)

	rank, err := strconv.Atoi(strings.Trim(parts[1], "\"()"))
	fmt.Println("rank", rank)
	if err != nil {
		fmt.Println("error", err)
		return errors.New("HouseImages Scan: invalid rank value")
	}

	// Set the values of the HouseImages struct
	hi.Url = url
	hi.Rank = rank

	return nil
}

func (hr *houseRepository) Get(id uuid.UUID) (*models.House, error) {
	fmt.Println("id", id)
	query := `SELECT h.id, h.user_id, h.kind, h.name, h.age, h.gender, h.content, l.id AS location_id, l.name AS location_name, c.id AS category_id, c.name AS category_name, h.created_at, ARRAY_AGG((hi.url, hi.rank)) AS house_images
				FROM houses AS h
				JOIN locations AS l ON h.location_id = l.id 
				JOIN categories As c ON h.category_id = c.id
				LEFT JOIN house_images AS hi ON h.id = hi.house_id
				WHERE h.id = $1
				GROUP BY h.id, l.id, c.id`

	row := hr.DB.QueryRow(query, id)

	var h models.House
	var houseImages []HouseImages
	err := row.Scan(&h.ID, &h.UserID, &h.Address, &h.IsRenting, &h.Price, &h.Size, &h.Kitchen, &h.Bathroom, &h.SleepingRoom, &h.CreatedAt, pq.Array(&houseImages))
	if err != nil {
		fmt.Println("error", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("houseImages", houseImages)
	fmt.Println("h", h)
	h.HouseImages = make([]models.HouseImage, len(houseImages))
	for i, hi := range houseImages {
		h.HouseImages[i] = models.HouseImage(hi)
	}
	return &h, nil
}

func (hr *houseRepository) Create(h *models.House) (uuid.UUID, error) {
	query := `
	INSERT INTO houses (id, user_id, address, is_renting, price, size, kitchen, bathroom, sleeping_room, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`
	_, err := hr.DB.Exec(
		query,
		h.ID,
		h.UserID,
		h.Address,
		h.IsRenting,
		h.Price,
		h.Size,
		h.Kitchen,
		h.Bathroom,
		h.SleepingRoom,
		h.CreatedAt,
	)
	if err != nil {
		return uuid.Nil, err
	}

	return h.ID, nil
}

func (hr *houseRepository) CreateHouseImage(id uuid.UUID, h models.HouseImage, houseID uuid.UUID) error {
	fmt.Println("houseID", houseID)
	query := `INSERT INTO house_images (id, url, rank, post_id) VALUES ($1, $2, $3, $4)`
	_, err := hr.DB.Exec(query, id, h.Url, h.Rank, houseID)
	if err != nil {
		return err
	}

	return nil
}

func (hr *houseRepository) DeleteAll() error {
	query := `DELETE FROM houses`
	_, err := hr.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
