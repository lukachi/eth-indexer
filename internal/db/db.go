package db

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn       *sql.DB
	SqlBuilder squirrel.StatementBuilderType
}

// ConnectDB initializes a standard *sql.DB and also sets up a squirrel builder.
func ConnectDB(user, password, host, port, dbname string) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Optionally, you might want to ping to check the connection:
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	// Initialize the squirrel builder with the connection placeholder (using dollar-sign placeholders for PostgreSQL).
	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	return &DB{
		Conn:       conn,
		SqlBuilder: builder,
	}, nil
}
