package postgresql

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/faruqisan/social-info/configs"
	"github.com/faruqisan/social-info/database"

	_ "github.com/lib/pq"
	"github.com/tokopedia/sqlt"
)

var databaseConnection string

// Postgresql .
type postgresql struct {
	db *sqlt.DB
}

// NewPostgresql .
func NewPostgresql() database.Database {

	cfg := configs.GetConfig()

	if !cfg.IsInitialized {
		configs.InitConfig()
	}

	dbConf := cfg.DBConfig

	databaseConnection = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Database)
	db, err := sqlt.Open("postgres", databaseConnection)
	if err != nil {
		log.Fatalln(err)
	}

	return &postgresql{
		db,
	}
}

// GetDatabase .
func (p *postgresql) GetDatabase() (*sqlt.DB, error) {

	if p.db == nil {
		return nil, errors.New("DBObject is nil")
	}

	dbStatuses, _ := p.db.GetStatus()

	for _, dbStatus := range dbStatuses {
		if !dbStatus.Connected {
			return nil, errors.New("DB not connected")
		}
	}

	return p.db, nil
}
