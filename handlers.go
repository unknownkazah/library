package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// Author
func CreateAuthor(c echo.Context) error {
	AuthorDB.Lock.Lock()
	defer AuthorDB.Lock.Unlock()

	u := &Author{}
	if err := c.Bind(u); err != nil { // 5
		return err
	}
	u.ID = AuthorDB.Sequence

	AuthorDB.Authors[u.ID] = u
	AuthorDB.Sequence++

	return c.JSON(http.StatusCreated, u)
}

func GetAuthor(c echo.Context) error {
	AuthorDB.Lock.Lock()
	defer AuthorDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, AuthorDB.Authors[id])
}

func GetAllAuthor(c echo.Context) error {
	AuthorDB.Lock.Lock()
	defer AuthorDB.Lock.Unlock()
	return c.JSON(http.StatusOK, AuthorDB.Authors)
}

func UpdateAuthor(c echo.Context) error {
	AuthorDB.Lock.Lock()
	defer AuthorDB.Lock.Unlock()

	u := new(Author)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	AuthorDB.Authors[id].Name = u.Name
	AuthorDB.Authors[id].Lastname = u.Lastname
	AuthorDB.Authors[id].Username = u.Username
	AuthorDB.Authors[id].Specialization = u.Specialization

	return c.JSON(http.StatusOK, AuthorDB.Authors[id])
}

func DeleteAuthor(c echo.Context) error {
	AuthorDB.Lock.Lock()
	defer AuthorDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(AuthorDB.Authors, id)
	return c.NoContent(http.StatusNoContent)
}

// Book
func CreateBook(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()

	u := &Book{}
	if err := c.Bind(u); err != nil { // 5
		return err
	}
	u.ID = BookDB.Sequence

	BookDB.Books[u.ID] = u
	BookDB.Sequence++

	return c.JSON(http.StatusCreated, u)
}

func GetBook(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, BookDB.Books[id])
}

func GetAllBooks(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()
	return c.JSON(http.StatusOK, BookDB.Books)
}

func UpdateBook(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()

	u := new(Book)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	BookDB.Books[id].Title = u.Title
	BookDB.Books[id].Genre = u.Genre
	BookDB.Books[id].CodeISBN = u.CodeISBN

	return c.JSON(http.StatusOK, BookDB.Books[id])

}

func DeleteBook(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(BookDB.Books, id)
	return c.NoContent(http.StatusNoContent)
}

//Member

func CreateMember(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()

	u := &Member{}
	if err := c.Bind(u); err != nil { // 5
		return err
	}
	u.ID = MemberDB.Sequence

	MemberDB.Members[u.ID] = u
	MemberDB.Sequence++

	return c.JSON(http.StatusCreated, u)
}

func GetMember(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, MemberDB.Members[id])
}

func GetAllMember(c echo.Context) error {
	BookDB.Lock.Lock()
	defer BookDB.Lock.Unlock()
	return c.JSON(http.StatusOK, MemberDB.Members)
}

func UpdateMember(c echo.Context) error {
	MemberDB.Lock.Lock()
	defer MemberDB.Lock.Unlock()

	u := new(Member)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	MemberDB.Members[id].Name = u.Name
	MemberDB.Members[id].Lastname = u.Lastname
	MemberDB.Members[id].BorrowedBooks = u.BorrowedBooks

	return c.JSON(http.StatusOK, MemberDB.Members[id])

}

func DeleteMember(c echo.Context) error {
	MemberDB.Lock.Lock()
	defer MemberDB.Lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(MemberDB.Members, id)
	return c.NoContent(http.StatusNoContent)
}
