package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// declare Postgree Database
type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// #region Method

// Postgree Connection Database
func ConnectPostgres(conn Postgres) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conn.Host,
		conn.Port,
		conn.User,
		conn.Password,
		conn.DBName,
		conn.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// #endregion Method
