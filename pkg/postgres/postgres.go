// PG driver implementation

package postgres

import (
	"fmt"
	"time"

	"golang.org/x/exp/slog"

	"database/sql"

	_ "github.com/lib/pq"
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

	db *sql.DB
}

// this function returns
// a new Postgres DB
func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("error occurred while opening the DataBase: %w", err)
	}
	err = db.Ping()
	if err != nil {
		slog.Info("PING: DB unavailable")
	} else {
		slog.Info("\033[32m:DB connection established \033[0m \n")
	}

	return db, err
}

// close active DB connection
func (db *Postgres) CloseDB() {
	if db.db != nil {
		db.db.Close()
	}
}
