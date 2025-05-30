// PG driver implementation

package postgres

import (
	"fmt"
	"log"
	"time"

	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

// default values
const (
	_defaultMaxPoolSize        = 1
	_defaultConnectionAttempts = 10
	_defaultConnTimeout        = 0
)

// Driver abstraction
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	db   *sql.DB
	Pool *pgxpool.Pool
}

// function to return
// a new instance of the
// Postgres Pool
func NewPool(url string) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnectionAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("error parsing DB-pool configs: %v", err)
	}
	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			break
		}
		log.Printf("Trying to connect to Postgres, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

// this function returns
// a new Postgres DB
func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("error occurred while opening the DataBase: %w", err)
	}
	err = db.Ping()

	return db, err
}

// close active connection pool
func (db *Postgres) ClosePool() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// close active DB connection
func (db *Postgres) CloseDB() {
	if db.db != nil {
		db.db.Close()
	}
}
