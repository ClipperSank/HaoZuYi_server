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

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) models.PostRepository {
	return &postRepository{
		DB: db,
	}
}

func (pr *postRepository) GetAll() ([]*models.Post, error) {
	query := `SELECT p.id, p.user_id, p.kind, p.name, p.age, p.gender, p.content, l.id AS location_id, l.name AS location_name, c.id AS category_id, c.name AS category_name, p.created_at
				FROM posts AS p
				JOIN locations AS l ON p.location_id = l.id 
				JOIN categories As c ON p.category_id = c.id 
				ORDER BY p.created_at DESC`

	fmt.Println("query", query)
	rows, err := pr.DB.Query(query)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Kind, &p.Name, &p.Age, &p.Gender, &p.Content, &p.Location.ID, &p.Location.Name, &p.Category.ID, &p.Category.Name, &p.CreatedAt)
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		posts = append(posts, &p)
	}
	fmt.Println("posts", posts)
	return posts, nil
}

type PostImages models.PostImages

func (pi *PostImages) Scan(src interface{}) error {
	fmt.Println("Scanning")
	if src == nil {
		return errors.New("PostImages Scan: null value")
	}

	// The src value returned by pq is of type []byte, so we need to convert it to a string
	srcStr, ok := src.([]byte)
	if !ok {
		return errors.New("PostImages Scan: invalid type")
	}
	fmt.Println("srcStr", srcStr)
	srcStr = bytes.Trim(srcStr, "{}")
	fmt.Println("srcStr", srcStr)
	// Split the string into URL and rank using the separator ","
	parts := strings.Split(string(srcStr), ",")
	if len(parts) != 2 {
		return errors.New("PostImages Scan: invalid format")
	}

	// Parse the URL and rank values from the string parts
	url := strings.Trim(parts[0], "\"()")
	fmt.Println("url", url)

	rank, err := strconv.Atoi(strings.Trim(parts[1], "\"()"))
	fmt.Println("rank", rank)
	if err != nil {
		fmt.Println("error", err)
		return errors.New("PostImages Scan: invalid rank value")
	}

	// Set the values of the PostImages struct
	pi.Url = url
	pi.Rank = rank

	return nil
}

func (pr *postRepository) Get(id uuid.UUID) (*models.Post, error) {
	fmt.Println("id", id)
	query := `SELECT p.id, p.user_id, p.kind, p.name, p.age, p.gender, p.content, l.id AS location_id, l.name AS location_name, c.id AS category_id, c.name AS category_name, p.created_at, ARRAY_AGG((pi.url, pi.rank)) AS post_images
				FROM posts AS p
				JOIN locations AS l ON p.location_id = l.id 
				JOIN categories As c ON p.category_id = c.id
				LEFT JOIN post_images AS pi ON p.id = pi.post_id
				WHERE p.id = $1
				GROUP BY p.id, l.id, c.id`

	row := pr.DB.QueryRow(query, id)

	var p models.Post
	var postImages []PostImages
	err := row.Scan(&p.ID, &p.UserID, &p.Kind, &p.Name, &p.Age, &p.Gender, &p.Content, &p.Location.ID, &p.Location.Name, &p.Category.ID, &p.Category.Name, &p.CreatedAt, pq.Array(&postImages))
	if err != nil {
		fmt.Println("error", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("postImages", postImages)
	fmt.Println("p", p)
	p.PostImages = make([]models.PostImages, len(postImages))
	for i, pi := range postImages {
		p.PostImages[i] = models.PostImages(pi)
	}
	return &p, nil
}

func (pr *postRepository) Create(p *models.Post) (uuid.UUID, error) {
	query := `INSERT INTO posts (id, user_id, kind, name, age, gender, content, location_id, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := pr.DB.Exec(query, p.ID, p.UserID, p.Kind, p.Name, p.Age, p.Gender, p.Content, p.Location.ID, p.Category.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return p.ID, nil
}

func (pr *postRepository) CreatePostImage(id uuid.UUID, p models.PostImages, postID uuid.UUID) error {
	fmt.Println("postID", postID)
	query := `INSERT INTO post_images (id, url, rank, post_id) VALUES ($1, $2, $3, $4)`
	_, err := pr.DB.Exec(query, id, p.Url, p.Rank, postID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *postRepository) DeleteAll() error {
	query := `DELETE FROM posts`
	_, err := pr.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
