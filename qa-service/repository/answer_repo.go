package repository

import (
	"database/sql"
	"time"
)

type Answer struct {
	ID         int64     `json:"id"`
	QuestionID int64     `json:"question_id"`
	UserID     int64     `json:"user_id"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}
type AnswerRepo struct {
	db *sql.DB
}

func NewAnswerRepo(db *sql.DB) *AnswerRepo {
	return &AnswerRepo{db: db}
}
