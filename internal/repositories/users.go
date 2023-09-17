package repositories

import (
	"database/sql"
	"fyno/server/internal/models"

	"github.com/google/uuid"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) models.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) Get(id uuid.UUID) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	row := u.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.IndexPage, &user.Username, &user.Role, &user.CreatedAt, &user.Account, &user.Password, &user.Age, &user.Birthday, &user.ContractCount, &user.HousesForRent, &user.OwnedHouses)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetByName(name string) (*models.User, error) {
	query := `SELECT * FROM users WHERE username = $1`
	row := u.DB.QueryRow(query, name)

	var user models.User
	err := row.Scan(&user.ID, &user.IndexPage, &user.Username, &user.Role, &user.CreatedAt, &user.Account, &user.Password, &user.Age, &user.Birthday, &user.ContractCount, &user.HousesForRent, &user.OwnedHouses)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (id, indexpage, username, role, created_at, account, password, age, birthday, contract_count, houses_for_rent, owned_houses) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
		RETURNING id`
	var id uuid.UUID
	err := u.DB.QueryRow(query, user.ID, user.IndexPage, user.Username, user.Role, user.CreatedAt, user.Account, user.Password, user.Age, user.Birthday, user.ContractCount, user.HousesForRent, user.OwnedHouses).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        id,
		IndexPage: user.IndexPage,
		Username:  user.Username,
		Role:      user.Role,
		// Include other fields here
	}, nil
}

func (u *userRepository) DeleteAll() error {
	query := `DELETE FROM users`
	_, err := u.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}