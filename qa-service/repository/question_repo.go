package repository

import (
	"context"
	"database/sql"
	"time"
)

type Question struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type QuestionRepo struct {
	db *sql.DB
}

func NewQuestionRepo(db *sql.DB) *QuestionRepo {
	return &QuestionRepo{db: db}
}

func (r *QuestionRepo) Create(ctx context.Context, q Question) error {
	query := `
		INSERT INTO questions (user_id, title, body, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, created_at
	`

	return r.db.QueryRow(query,
		q.UserID, q.Title, q.Body,
	).Scan(&q.ID, &q.CreatedAt)
}

func (r *QuestionRepo) GetByID(ctx context.Context, id int64) (*Question, error) {
	query := `
		SELECT id, user_id, title, body, created_at
		FROM questions WHERE id=$1
	`

	var q Question
	err := r.db.QueryRow(query, id).
		Scan(&q.ID, &q.UserID, &q.Title, &q.Body, &q.CreatedAt)

	return &q, err
}
