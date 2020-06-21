package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type row struct{}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == production
}

func Open(driverName, dataSourceName string) (SqlClient, error) {
	if isMocked && !isProduction() {
		dbClient = &clientMock{}
		return dbClient, nil
	}
	if driverName == "" {
		return nil, errors.New("Invalid driver name")
	}

	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	dbClient = &client{
		db: database,
	}

	return dbClient, nil
}

func (c *client) Query(query string, args ...interface{}) (rows, error) {
	returnedRows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	result := sqlRows{
		rows: returnedRows,
	}
	return &result, nil
}
