package member

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Create(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	m := &member{}
	if err = c.Bind(m); err != nil { // 5
		return err
	}
	m.ID = database.Sequence

	database.Map[m.ID] = m
	database.Sequence++

	return c.JSON(http.StatusCreated, m)
}

func Get(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, database.Map[id])
}

func GetAll(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()
	return c.JSON(http.StatusOK, database.Map)
}

func Update(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	m := new(member)
	if err = c.Bind(m); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	database.Map[id].Name = m.Name
	database.Map[id].Lastname = m.Lastname
	database.Map[id].BorrowedBooks = m.BorrowedBooks

	return c.JSON(http.StatusOK, database.Map[id])

}

func Delete(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	delete(database.Map, id)

	return c.NoContent(http.StatusNoContent)
}
