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
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetByName(name string) (*models.User, error) {
	query := `SELECT * FROM users WHERE name = $1`
	row := u.DB.QueryRow(query, name)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Create(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (id, name, email, avatar_url) VALUES ($1, $2, $3, $4) RETURNING id`
	var id uuid.UUID
	err := u.DB.QueryRow(query, user.ID, user.Name, user.Email, user.AvatarURL).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
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
