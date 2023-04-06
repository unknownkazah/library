package member

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Create(c echo.Context) (err error) {
	data := Member{}
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
	data := Member{}
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
