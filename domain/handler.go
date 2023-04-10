package main

import (
	"bibliotekaProject/domain/author"
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)


type Crud interface {
	Create(c echo.Context)(err error)
	Get(c echo.Context) (err error)
	GetAll(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)

}

type Handler struct {
	storage author.Storage
}

func NewHandler(storage author.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Create(c echo.Context) (err error) {
	data := author.Author{}
	if err = c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	data.ID, err = h.storage.CreateRow(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, data)
}

func (h *Handler) Get(c echo.Context) (err error) {
	id := c.Param("id")
	data, err := h.storage.GetRowByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) GetAll(c echo.Context) (err error) {
	data, err := h.storage.SelectRows()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) Update(c echo.Context) (err error) {
	id := c.Param("id")

	_, err = h.storage.GetRowByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	data := author.Author{}
	if err = c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data.ID = c.Param("id")

	err = h.storage.UpdateRow(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) Delete(c echo.Context) (err error) {
	id := c.Param("id")

	err = h.storage.DeleteRow(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

//var books []book.Book
//
//func (h *Handler) GetBooksByAuthorID(c echo.Context) (err error) {
//	id := c.Param("id")
//	data, err := h.storage.GetRowByID(id)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return c.String(http.StatusNotFound, err.Error())
//		}
//		return c.String(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, data)
	//
	//authorBooks := []book.Book{}
	//for _, bookid := range books {
	//	if bookid.ID == id {
	//		authorBooks = append(authorBooks, bookid)
	//	}
	//}
	//_, err = h.storage.GetRowByID(id)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return c.String(http.StatusNotFound, err.Error())
	//	}
	//	return c.String(http.StatusInternalServerError, err.Error())
	//}
	//return c.JSON(http.StatusOK, authorBooks)
}