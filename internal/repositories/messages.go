package repositories

import (
	"database/sql"
	"fyno/server/internal/models"
	"time"

	"github.com/google/uuid"
)

type messageRepository struct {
	DB *sql.DB
}

func NewMessageRepository(db *sql.DB) models.MessageRepository {
	return &messageRepository{
		DB: db,
	}
}

func (m *messageRepository) GetAll(userID uuid.UUID) ([]models.Message, error) {
	query := `SELECT id, sender, receiver, content, created_at FROM messages WHERE sender=$1 OR receiver=$1 ORDER BY created_at ASC`
	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.Sender, &msg.Receiver, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func (m *messageRepository) Create(msg *models.Message) (uuid.UUID, error) {
	msg.ID = uuid.New()
	msg.CreatedAt = time.Now()

	query := `INSERT INTO messages (id, sender, receiver, content, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := m.DB.Exec(query, msg.ID, msg.Sender, msg.Receiver, msg.Content, msg.CreatedAt)
	if err != nil {
		return uuid.Nil, err
	}

	return msg.ID, nil
}

// Get all users from the message_user_groups table ordered by updated_at
func (m *messageRepository) GetAllUserGroups(_userID uuid.UUID) ([]models.User, error) {
	query := `
		SELECT u.id, u.username, u.role, u.created_at, u.account, u.password, u.age, u.birthday, u.contract_count, u.houses_for_rent, u.owned_houses
		FROM message_user_groups m
		INNER JOIN users u ON (m.user_id = u.id OR m.message_partner_id = u.id)
		WHERE (m.user_id = $1 OR m.message_partner_id = $1)
		ORDER BY m.updated_at DESC
	`

	rows, err := m.DB.Query(query, _userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList []models.User
	// Use a map to keep track of users that have already been added to the list
	usersMap := make(map[uuid.UUID]struct{})
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
			&user.Account,
			&user.Password,
			&user.Age,
			&user.Birthday,
			&user.ContractCount,
			&user.HousesForRent,
			&user.OwnedHouses,
		)
		if err != nil {
			return nil, err
		}
		// If the user has not been added to the list yet, add them
		if _, ok := usersMap[user.ID]; !ok {
			usersList = append(usersList, user)
			usersMap[user.ID] = struct{}{}
		}
	}

	return usersList, nil
}


func (m *messageRepository) CreateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error {
	query := `INSERT INTO message_user_groups (id, user_id, message_partner_id) VALUES ($1, $2, $3)`
	_, err := m.DB.Exec(query, uuid.New(), userID, messagePartnerID)
	if err != nil {
		return err
	}

	return nil
}

func (m *messageRepository) UpdateUserGroup(userID uuid.UUID, messagePartnerID uuid.UUID) error {
	query := `UPDATE message_user_groups SET updated_at = $1 WHERE user_id = $2 AND message_partner_id = $3`
	_, err := m.DB.Exec(query, time.Now(), userID, messagePartnerID)
	if err != nil {
		return err
	}

	return nil
}

func (m *messageRepository) IsUserGroupExists(userID uuid.UUID, messagePartnerID uuid.UUID) (bool, error) {
	query := `SELECT id FROM message_user_groups WHERE user_id = $1 AND message_partner_id = $2`
	var id uuid.UUID
	err := m.DB.QueryRow(query, userID, messagePartnerID).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
