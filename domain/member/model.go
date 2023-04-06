package member

type Member struct {
	ID            string  `json:"id" db:"id"`
	Name          *string `json:"name" db:"name"`
	Lastname      *string `json:"lastname" db:"lastname"`
	BorrowedBooks *string `json:"borrowedBooks" db:"borrowed_books"`
}
