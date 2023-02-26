package member

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Create(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	u := &member{}
	if err := c.Bind(u); err != nil { // 5
		return err
	}
	u.ID = database.Sequence

	database.Members[u.ID] = u
	database.Sequence++

	return c.JSON(http.StatusCreated, u)
}

func Get(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, database.Members[id])
}

func GetAll(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()
	return c.JSON(http.StatusOK, database.Members)
}

func Update(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	u := new(member)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	database.Members[id].Name = u.Name
	database.Members[id].Lastname = u.Lastname
	database.Members[id].BorrowedBooks = u.BorrowedBooks

	return c.JSON(http.StatusOK, database.Members[id])

}

func Delete(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	delete(database.Members, id)

	return c.NoContent(http.StatusNoContent)
}
