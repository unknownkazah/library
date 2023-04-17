package dto

import "bibliotekaProject/internal/entity"

type MemberRequest struct {
	ID            string `db:"id"`
	Name          string `db:"name"`
	Lastname      string `db:"lastname"`
	BorrowedBooks string `db:"borrowed_books"`
	MemberIdBooks string `db:"member_id_books"`
}

type MemberResponse struct {
	ID            string `db:"id"`
	Name          string `db:"name"`
	Lastname      string `db:"lastname"`
	BorrowedBooks string `db:"borrowed_books"`
	MemberIdBooks string `db:"member_id_books"`
}

func ParseFromMember(src entity.Member) (dst MemberResponse) {
	dst = MemberResponse{
		ID:            src.ID,
		Name:          *src.Name,
		Lastname:      *src.Lastname,
		BorrowedBooks: *src.BorrowedBooks,
		MemberIdBooks: *src.MemberIdBooks,
	}

	return
}

func ParseFromMembers(src []entity.Member) (dst []MemberResponse) {
	for _, data := range src {
		dst = append(dst, ParseFromMember(data))
	}

	return
}
