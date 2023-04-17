package entity

type Member struct {
	ID            string  `db:"id"`
	Name          *string `db:"name"`
	Lastname      *string `db:"lastname"`
	BorrowedBooks *string `db:"borrowed_books"`
}
