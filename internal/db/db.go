package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sqlx.DB
}

func ConnectDB(user, password, host, port, dbname string) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{Conn: conn}, nil
}

func (db *DB) CreateTables() error {
	blockTable := `
      CREATE TABLE IF NOT EXISTS blocks (
          number BIGINT PRIMARY KEY,
          hash TEXT,
          parent_hash TEXT,
          timestamp BIGINT
      );`
	txTable := `
      CREATE TABLE IF NOT EXISTS transactions (
          hash TEXT PRIMARY KEY,
          "from" TEXT,
          "to" TEXT,
          value TEXT,
          block_number BIGINT REFERENCES blocks(number)
      );`
	if _, err := db.Conn.Exec(blockTable); err != nil {
		return err
	}
	if _, err := db.Conn.Exec(txTable); err != nil {
		return err
	}
	return nil
}
