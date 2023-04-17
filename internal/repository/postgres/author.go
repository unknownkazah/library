package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"

	"bibliotekaProject/internal/entity"
)

type AuthorRepository struct {
	db *sqlx.DB
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (s *AuthorRepository) CreateRow(data entity.Author) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO authors (name, lastname, username, specialization)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	args := []any{data.Name, data.Lastname, data.Username, data.Specialization}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)

	return
}

func (s *AuthorRepository) GetRowByID(id string) (dest entity.Author, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, name, lastname, username, specialization
		FROM authors
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *AuthorRepository) SelectRows() (dest []entity.Author, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, name, lastname, username, specialization
		FROM authors
		ORDER BY created_at`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *AuthorRepository) DeleteRow(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM authors
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}

func (s *AuthorRepository) UpdateRow(data entity.Author) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, data.ID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE authors SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}

	return
}

func (s *AuthorRepository) prepareArgs(data entity.Author) (sets []string, args []any) {
	if data.Username != nil {
		args = append(args, data.Username)
		sets = append(sets, fmt.Sprintf("username=$%d", len(args)))

	}

	if data.Lastname != nil {
		args = append(args, data.Lastname)
		sets = append(sets, fmt.Sprintf("lastname=$%d", len(args)))
	}

	if data.Name != nil {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}

	if data.Specialization != nil {
		args = append(args, data.Specialization)
		sets = append(sets, fmt.Sprintf("specialization=$%d", len(args)))
	}

	return
}
