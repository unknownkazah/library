package health

import (
	"context"
	"net/http"
	"time"

	"github.com/hellofresh/health-go/v5"
	httpHealth "github.com/hellofresh/health-go/v5/checks/http"
	postgresHealth "github.com/hellofresh/health-go/v5/checks/postgres"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	googleURL   string
	postgresDSN string
}

func NewHandler(googleURL, postgresDSN string) *Handler {
	return &Handler{
		googleURL:   googleURL,
		postgresDSN: postgresDSN,
	}
}

func (h *Handler) Healthcheck(c echo.Context) (err error) {
	check, err := health.New(
		health.WithSystemInfo(),
		health.WithComponent(health.Component{
			Name:    "library",
			Version: "1.0.0"}),
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// http health check example
	err = check.Register(health.Config{
		Name:      "google",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: httpHealth.New(httpHealth.Config{
			URL: h.googleURL,
		}),
	})

	// postgres health check example
	//
	err = check.Register(health.Config{
		Name:      "postgres",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: postgresHealth.New(postgresHealth.Config{
			DSN: h.postgresDSN,
		}),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, check.Measure(context.Background()))
}
