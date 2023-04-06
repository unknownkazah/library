package database

import (
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// postgres://username:password@localhost:5432/dbname?sslmode=disable

// New established connection to a database instance using provided URI and auth credentials.
func New(databaseSourceName string) (db *sqlx.DB, err error) {
	if !strings.Contains(databaseSourceName, "://") {
		err = errors.New("sql: unknown data source name " + databaseSourceName)
		return
	}
	driverName := strings.Split(databaseSourceName, "://")[0]

	db, err = sqlx.Connect(driverName, databaseSourceName)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20)

	switch driverName {
	case "postgres":
		_, err = db.Exec(`SET TIMEZONE='Asia/Almaty';`)
		if err != nil {
			return
		}
	}

	return
}
