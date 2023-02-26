package main

import (
	"bibliotekaProject/author"
	"bibliotekaProject/book"
	"bibliotekaProject/member"
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

	// Authors
	authors := api.Group("/authors")
	authors.POST("", author.Create)
	authors.GET("/:id", author.Get)
	authors.GET("", author.GetAll)
	authors.PUT("/:id", author.Update)
	authors.DELETE("/:id", author.Delete)

	// Books
	books := api.Group("books")
	books.POST("", book.Create)
	books.GET("/:id", book.Get)
	books.GET("", book.GetAll)
	books.PUT("/:id", book.Update)
	books.DELETE("/:id", book.Delete)

	// Members
	members := api.Group("/members")
	members.POST("", member.Create)
	members.GET("/:id", member.Get)
	members.GET("", member.GetAll)
	members.PUT("/:id", member.Update)
	members.DELETE("/:id", member.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
