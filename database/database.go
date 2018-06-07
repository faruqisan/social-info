package database

import "github.com/tokopedia/sqlt"

// Database interface contracts
type Database interface {
	GetDatabase() (*sqlt.DB, error)
}
