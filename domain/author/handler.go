package author


type Handler struct {
	storage Storage
}
func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
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