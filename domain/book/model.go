package book

import "sync"

type book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	CodeISBN int    `json:"codeISBN"`
}

type table struct {
	Map      map[int]*book
	Sequence int
	sync.Mutex
}

var database = table{
	Map:      map[int]*book{},
	Sequence: 1,
	Mutex:    sync.Mutex{},
}
