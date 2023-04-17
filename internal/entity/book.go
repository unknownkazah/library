package entity

type Book struct {
	ID            string  `db:"id"`
	Title         *string `db:"title"`
	Genre         *string `db:"genre"`
	CodeISBN      *string `db:"code_isbn"`
	IdAuthorBooks *string `db:"id_author_books"`
}
