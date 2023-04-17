package dto

import "bibliotekaProject/internal/entity"

type AuthorRequest struct {
	ID             string
	Name           string `json:"name" validate:"required" `
	Lastname       string `json:"lastname" validate:"required" `
	Username       string `json:"username" validate:"required" `
	Specialization string `json:"specialization" validate:"required" `
}

type AuthorResponse struct {
	ID             string `json:"ID"`
	Name           string `json:"name" `
	Lastname       string `json:"lastname" `
	Username       string `json:"username" `
	Specialization string `json:"specialization" `
}

func ParseFromAuthor(src entity.Author) (dst AuthorResponse) {
	dst = AuthorResponse{
		ID:             src.ID,
		Name:           *src.Name,
		Lastname:       *src.Lastname,
		Username:       *src.Username,
		Specialization: *src.Specialization,
	}

	return
}

func ParseFromAuthors(src []entity.Author) (dst []AuthorResponse) {
	for _, data := range src {
		dst = append(dst, ParseFromAuthor(data))
	}

	return
}
