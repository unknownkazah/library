package library

import (
	"bibliotekaProject/internal/dto"
	"bibliotekaProject/internal/entity"
)

func (s *Service) CreateAuthor(req dto.AuthorRequest) (res dto.AuthorResponse, err error) {
	author := entity.Author{
		Name:           &req.Name,
		Lastname:       &req.Lastname,
		Username:       &req.Username,
		Specialization: &req.Specialization,
	}

	author.ID, err = s.authors.CreateRow(author)
	if err != nil {
		return
	}
	res = dto.ParseFromAuthor(author)

	return
}

func (s *Service) GetAuthor(id string) (res dto.AuthorResponse, err error) {
	author, err := s.authors.GetRowByID(id)
	if err != nil {
		return
	}
	res = dto.ParseFromAuthor(author)

	return
}

func (s *Service) GetAuthors() (res []dto.AuthorResponse, err error) {
	authors, err := s.authors.SelectRows()
	if err != nil {
		return
	}
	res = dto.ParseFromAuthors(authors)

	return
}

func (s *Service) GetBooksByAuthorID(id string) (res dto.BookResponse, err error) {
	book, err := s.books.SelectBookByAuthorID(id)
	if err != nil {
		return
	}
	res = dto.ParseFromBook(book)

	return

}

//aa
func (s *Service) UpdateAuthor(req dto.AuthorRequest) (err error) {
	author := entity.Author{
		Name:           &req.Name,
		Lastname:       &req.Lastname,
		Username:       &req.Username,
		Specialization: &req.Specialization,
	}

	return s.authors.UpdateRow(author)
}

func (s *Service) DeleteAuthor(id string) (err error) {
	return s.authors.DeleteRow(id)
}
