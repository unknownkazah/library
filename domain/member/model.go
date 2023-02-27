package member

import "sync"

type member struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	BorrowedBooks string `json:"borrowedBooks"`
}

type table struct {
	Map      map[int]*member
	Sequence int
	sync.Mutex
}

var database = table{
	Map:      map[int]*member{},
	Sequence: 1,
	Mutex:    sync.Mutex{},
}
