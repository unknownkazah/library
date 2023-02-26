package author

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Create(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	a := &author{}
	if err = c.Bind(a); err != nil {
		return
	}
	a.ID = database.Index

	database.Map[a.ID] = a
	database.Index++

	return c.JSON(http.StatusCreated, a)
}

func Get(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

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

	a := new(author)
	if err = c.Bind(a); err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	database.Map[id].Name = a.Name
	database.Map[id].Lastname = a.Lastname
	database.Map[id].Username = a.Username
	database.Map[id].Specialization = a.Specialization

	return c.JSON(http.StatusOK, database.Map[id])
}

func Delete(c echo.Context) (err error) {
	database.Lock()
	defer database.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	delete(database.Map, id)

	return c.NoContent(http.StatusNoContent)
}
