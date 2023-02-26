package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	api := e.Group("/api")

	// Routes
	authors := api.Group("/authors")
	authors.POST("", CreateAuthor)
	authors.GET("/:id", GetAuthor)
	authors.GET("", GetAllAuthor)
	authors.PUT("/:id", UpdateAuthor)
	authors.DELETE("/:id", DeleteAuthor)

	books := api.Group("books")
	books.POST("", CreateBook)
	books.GET("/:id", GetBook)
	books.GET("", GetAllBooks)
	books.PUT("/:id", UpdateBook)
	books.DELETE("/:id", DeleteBook)

	members := api.Group("/members")
	members.POST("", CreateMember)
	members.GET("/:id", GetMember)
	members.GET("", GetAllMember)
	members.PUT("/:id", UpdateMember)
	members.DELETE("/:id", DeleteMember)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
