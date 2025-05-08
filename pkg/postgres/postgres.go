// PG driver implementation

package postgres

import (
	"fmt"
	"log"
	"time"

	"context"

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

	Pool *pgxpool.Pool
}

// function to return
// a new instance of the
// Postgres
func New(url string) (*Postgres, error) {
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

// close active connection
func (db *Postgres) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
