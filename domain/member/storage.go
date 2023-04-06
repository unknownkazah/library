package member

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type Storage interface {
	CreateRow(data Member) (dest string, err error)
	GetRowByID(id string) (dest Member, err error)
	DeleteRow(id string) (err error)
	SelectRows() (dest []Member, err error)
	UpdateRow(data Member) (err error)
}

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreateRow(data Member) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO members (name, lastname, borrowed_books)
		VALUES ($1, $2, $3)
		RETURNING id`

	args := []any{data.Name, data.Lastname, data.BorrowedBooks}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)

	return
}

func (s *storage) GetRowByID(id string) (dest Member, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, name, lastname, borrowed_books
		FROM members
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *storage) SelectRows() (dest []Member, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id, name, lastname, borrowed_books
		FROM members
		ORDER BY created_at`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *storage) DeleteRow(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE 
		FROM books
		WHERE id=$1`

	args := []any{id}

	_, err = s.db.ExecContext(ctx, query, args...)

	return
}

func (s *storage) UpdateRow(data Member) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sets, args := s.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, data.ID)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")

		query := fmt.Sprintf("UPDATE members SET %s WHERE id=$%d", strings.Join(sets, ", "), len(args))
		_, err = s.db.ExecContext(ctx, query, args...)
	}

	return
}

func (s *storage) prepareArgs(data Member) (sets []string, args []any) {

	if data.Lastname != nil {
		args = append(args, data.Lastname)
		sets = append(sets, fmt.Sprintf("lastname=$%d", len(args)))
	}

	if data.Name != nil {
		args = append(args, data.Name)
		sets = append(sets, fmt.Sprintf("name=$%d", len(args)))
	}

	if data.BorrowedBooks != nil {
		args = append(args, data.BorrowedBooks)
		sets = append(sets, fmt.Sprintf("specialization=$%d", len(args)))
	}

	return
}
