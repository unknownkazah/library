package postgres

import (
	"bibliotekaProject/internal/entity"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type MemberRepository struct {
	db *sqlx.DB
}

func NewMemberRepository(db *sqlx.DB) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (s *MemberRepository) CreateRow(data entity.Member) (dest string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO members (name, lastname, borrowed_books, member_id_books)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	args := []any{data.Name, data.Lastname, data.BorrowedBooks, data.MemberIdBooks}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(&dest)

	return
}

func (s *MemberRepository) GetRowByID(id string) (dest entity.Member, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id,member_id_books, name, lastname, borrowed_books
		FROM members
		WHERE id=$1`

	args := []any{id}

	err = s.db.GetContext(ctx, &dest, query, args...)

	return
}

func (s *MemberRepository) SelectRows() (dest []entity.Member, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT id,member_id_books, name, lastname, borrowed_books
		FROM members
		ORDER BY created_at`

	err = s.db.SelectContext(ctx, &dest, query)

	return
}

func (s *MemberRepository) SelectMemberIdBooks(MemberID string) (dest []entity.Member, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT member_id_books, id , name, lastname, borrowed_books
		FROM members
		WHERE member_id_books=$1`

	args := []any{MemberID}

	err = s.db.SelectContext(ctx, &dest, query, args...)

	return
}

func (s *MemberRepository) DeleteRow(id string) (err error) {
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

func (s *MemberRepository) UpdateRow(data entity.Member) (err error) {
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

func (s *MemberRepository) prepareArgs(data entity.Member) (sets []string, args []any) {

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
