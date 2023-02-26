package main

import (
	"sync"
)

var (
	AuthorDB = AuthorMap{
		Authors:  map[int]*Author{},
		Sequence: 1,
		Lock:     sync.Mutex{},
	}
	BookDB = BookMap{
		Books:    map[int]*Book{},
		Sequence: 1,
		Lock:     sync.Mutex{},
	}
	MemberDB = MemberMap{
		Members:  map[int]*Member{},
		Sequence: 1,
		Lock:     sync.Mutex{},
	}
)

type Author struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Username       string `json:"username"`
	Specialization string `json:"specialization"`
}

type AuthorMap struct {
	Authors  map[int]*Author
	Sequence int
	Lock     sync.Mutex
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	CodeISBN int    `json:"codeISBN"`
}

type BookMap struct {
	Books    map[int]*Book
	Sequence int
	Lock     sync.Mutex
}

type Member struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	BorrowedBooks string `json:"borrowedBooks"`
}

type MemberMap struct {
	Members  map[int]*Member
	Sequence int
	Lock     sync.Mutex
}
