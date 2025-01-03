package db

import (
	"context"
	"database/sql"
	"time"
)

// NewConnection establishes a connection to a PostgreSQL database and returns a database connection pool.
func NewConnection(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {

	// Open a new database connection using the provided PostgreSQL address
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of idle connections to the database. These connections are not in use but are kept open for future use.
	db.SetMaxIdleConns(maxIdleConns)

	// Set the maximum number of open connections to the database. This controls the number of concurrent active connections.
	db.SetMaxOpenConns(maxIdleConns)

	// Parse the maxIdleTime string into a time.Duration (like 5s, 1m, etc.). This sets how long a connection can be idle before it is closed.
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	// Create a new context with a timeout of 5 seconds. This context will be used to check the connection.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to make sure the connection is working and the database is reachable.
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
