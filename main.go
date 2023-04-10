package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bibliotekaProject/domain/author"
	"bibliotekaProject/domain/book"
	"bibliotekaProject/domain/member"
	"bibliotekaProject/health"
	"bibliotekaProject/pkg/database"
)

func main() {
	googleURL := os.Getenv("GOOGLE_URL")
	postgresDSN := os.Getenv("POSTGRES_DSN")

	// create connection to database
	//
	//
	postgres, err := database.New(postgresDSN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer postgres.Close()

	// create migration to database
	err = database.Migrate(postgresDSN)
	if err != nil {
		fmt.Println(err)
		return
	}

	authorStorage := author.NewStorage(postgres)
	authorHandler := author.NewHandler(authorStorage)

	bookStorage := book.NewStorage(postgres)
	bookHandler := book.NewHandler(bookStorage)

	memberStorage := member.NewStorage(postgres)
	memberHandler := member.NewHandler(memberStorage)

	healthHandler := health.NewHandler(googleURL, postgresDSN)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	apiGroup := e.Group("/api")

	// Authors
	authorsGroup := apiGroup.Group("/authors")
	authorsGroup.POST("", authorHandler.Create)
	authorsGroup.GET("/:id", authorHandler.Get)
	//qauthorsGroup.GET("/:id/book", authorHandler.GetBooksByAuthorID)
	authorsGroup.GET("", authorHandler.GetAll)
	authorsGroup.PUT("/:id", authorHandler.Update)
	authorsGroup.DELETE("/:id", authorHandler.Delete)

	// Books
	booksGroup := apiGroup.Group("/books")
	booksGroup.POST("", bookHandler.Create)
	booksGroup.GET("/:id", bookHandler.Get)
	booksGroup.GET("", bookHandler.GetAll)
	booksGroup.PUT("/:id", bookHandler.Update)
	booksGroup.DELETE("/:id", bookHandler.Delete)

	// Members
	membersGroup := apiGroup.Group("/members")
	membersGroup.POST("", memberHandler.Create)
	membersGroup.GET("/:id", memberHandler.Get)
	membersGroup.GET("", memberHandler.GetAll)
	membersGroup.PUT("/:id", memberHandler.Update)
	membersGroup.DELETE("/:id", memberHandler.Delete)

	// Health
	healthGroup := apiGroup.Group("/health")
	healthGroup.GET("", healthHandler.Healthcheck)

	// Start server//
	e.Logger.Fatal(e.Start(":8080"))
}
