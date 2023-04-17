package dto

import "bibliotekaProject/internal/entity"

type BookRequest struct {
	ID            string `db:"id" `
	Title         string `db:"title" validate:"required"`
	Genre         string `db:"genre" validate:"required"`
	CodeISBN      string `db:"code_isbn" validate:"required"`
	IdAuthorBooks string `db:"id_author_books" validate:"required"`
}
type BookResponse struct {
	IdAuthorBooks string `db:"id_author_books"`
	ID            string `db:"id"`
	Title         string `db:"title"`
	Genre         string `db:"genre"`
	CodeISBN      string `db:"code_isbn"`
}

func ParseFromBook(src entity.Book) (dst BookResponse) {
	dst = BookResponse{
		ID:            src.ID,
		Title:         *src.Title,
		Genre:         *src.Genre,
		CodeISBN:      *src.CodeISBN,
		IdAuthorBooks: *src.IdAuthorBooks,
	}

	return
}

func ParseFromBooks(src []entity.Book) (dst []BookResponse) {
	for _, data := range src {
		dst = append(dst, ParseFromBook(data))
	}

	return
}
