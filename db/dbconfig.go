package db

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/services/db-models"
	"github.com/services/utils/config"
	"go.uber.org/zap"
)

var (
	once    sync.Once
	queries *db.Queries
	pool    *sql.DB
	err     error
)

func Connect_Database() {
	logger := config.GetLogger()
	DATABASE_URL := os.Getenv("POSTGRES_URL")
	once.Do(func() {
		pool, err = sql.Open("postgres", DATABASE_URL)
		if err != nil {
			logger.Error("Failed to connect to database: ", zap.Error(err))
		}

		if err = pool.Ping(); err != nil {
			logger.Error("Database failed to respond: ", zap.Error(err))
			return
		}
		queries = db.New(pool)
	})
}
func GetQuery() *db.Queries {
	return queries
}
