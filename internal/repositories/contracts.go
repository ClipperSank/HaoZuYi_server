package repositories

import (
	"database/sql"
	"fmt"
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type contractRepository struct {
	DB *sql.DB
}

func NewContractRepository(db *sql.DB) models.ContractRepository {
	return &contractRepository{
		DB: db,
	}
}

func (cr *contractRepository) GetAll() ([]*models.Contract, error) {
    query := `SELECT id, renter_id, landlord_id, house_id, contract, rent, start_time, end_time, renter_review, landlord_review
              FROM contracts
              ORDER BY renter_review DESC NULLS LAST`

    fmt.Println("query", query)
    rows, err := cr.DB.Query(query)
    if err != nil {
        fmt.Println("error", err)
        return nil, err
    }
    defer rows.Close()

    var contracts []*models.Contract
    for rows.Next() {
        var c models.Contract
        err := rows.Scan(&c.ID, &c.RenterID, &c.LandlordID, &c.HouseID, &c.ContractText, &c.Rent, &c.StartTime, &c.EndTime, &c.RenterReview, &c.LandlordReview)
        if err != nil {
            fmt.Println("error", err)
            return nil, err
        }
        contracts = append(contracts, &c)
    }
    fmt.Println("contracts", contracts)
    return contracts, nil
}

func (cr *contractRepository) Get(id uuid.UUID) (*models.Contract, error) {
    query := `SELECT id, renter_id, landlord_id, house_id, contract, rent, start_time, end_time, renter_review, landlord_review
              FROM contracts
              WHERE id = $1`

    row := cr.DB.QueryRow(query, id)

    var c models.Contract
    err := row.Scan(&c.ID, &c.RenterID, &c.LandlordID, &c.HouseID, &c.ContractText, &c.Rent, &c.StartTime, &c.EndTime, &c.RenterReview, &c.LandlordReview)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // Return nil if the contract with the given ID doesn't exist
        }
        return nil, err
    }

    return &c, nil
}

func (cr *contractRepository) Create(c *models.Contract) (uuid.UUID, error) {
	query := `
		INSERT INTO contracts (id, renter_id, landlord_id, house_id, contract, rent, start_time, end_time, renter_review, landlord_review)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := cr.DB.Exec(
		query,
		c.ID,
		c.RenterID,
		c.LandlordID,
		c.HouseID,
		c.ContractText,
		c.Rent,
		c.StartTime,
		c.EndTime,
		c.RenterReview,
		c.LandlordReview,
	)
	if err != nil {
		return uuid.Nil, err
	}

	return c.ID, nil
}

func (cr *contractRepository) DeleteAll() error {
	query := `DELETE FROM contracts`
	_, err := cr.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
