package library

import (
	"bibliotekaProject/internal/dto"
	"bibliotekaProject/internal/entity"
)

func (s *Service) CreateBook(req dto.BookRequest) (res dto.BookResponse, err error) {
	data := entity.Book{
		Title:         &req.Title,
		Genre:         &req.Genre,
		CodeISBN:      &req.CodeISBN,
		IdAuthorBooks: &req.IdAuthorBooks,
	}

	data.ID, err = s.books.CreateRow(data)
	if err != nil {
		return
	}
	res = dto.ParseFromBook(data)

	return
}

func (s *Service) GetBook(id string) (res dto.BookResponse, err error) {
	data, err := s.books.GetRowByID(id)
	if err != nil {
		return
	}
	res = dto.ParseFromBook(data)

	return
}

func (s *Service) GetBooks() (res []dto.BookResponse, err error) {
	data, err := s.books.SelectRows()
	if err != nil {
		return
	}
	res = dto.ParseFromBooks(data)

	return
}

func (s *Service) UpdateBook(req dto.BookRequest) (err error) {
	data := entity.Book{
		Title:         &req.Title,
		Genre:         &req.Genre,
		CodeISBN:      &req.CodeISBN,
		IdAuthorBooks: &req.IdAuthorBooks,
	}

	return s.books.UpdateRow(data)
}

func (s *Service) DeleteBook(id string) (err error) {
	return s.books.DeleteRow(id)
}
