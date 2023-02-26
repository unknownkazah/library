package author

import "sync"

type author struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Username       string `json:"username"`
	Specialization string `json:"specialization"`
}

type table struct {
	Map   map[int]*author
	Index int
	sync.Mutex
}

var database = table{
	Map:   map[int]*author{},
	Index: 1,
	Mutex: sync.Mutex{},
}
