package book

type Book struct {
	ID       string  `json:"id" db:"id"`
	Title    *string `json:"title" db:"title"`
	Genre    *string `json:"genre" db:"genre"`
	CodeISBN *string `json:"codeISBN" db:"code_isbn"`
}
