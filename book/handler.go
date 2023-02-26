package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Create(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	u := &book{}
	if err := c.Bind(u); err != nil { // 5
		return err
	}
	u.ID = database.Sequence

	database.Map[u.ID] = u
	database.Sequence++

	return c.JSON(http.StatusCreated, u)
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

	b := new(book)
	if err = c.Bind(b); err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	database.Map[id].Title = b.Title
	database.Map[id].Genre = b.Genre
	database.Map[id].CodeISBN = b.CodeISBN

	return c.JSON(http.StatusOK, database.Map[id])

}

func Delete(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	delete(database.Map, id)

	return c.NoContent(http.StatusNoContent)
}