package entity

type Author struct {
	ID             string  `db:"id"`
	Name           *string `db:"name"`
	Lastname       *string `db:"lastname"`
	Username       *string `db:"username"`
	Specialization *string `db:"specialization"`
}
