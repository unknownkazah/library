package author

type Author struct {
	ID             string  `json:"id" db:"id"`
	Name           *string `json:"name" db:"name"`
	Lastname       *string `json:"lastname" db:"lastname"`
	Username       *string `json:"username" db:"username"`
	Specialization *string `json:"specialization" db:"specialization"`
}
