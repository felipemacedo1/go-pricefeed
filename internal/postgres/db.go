package postgres

import (
	"database/sql"
	"fmt"

	"github.com/growthfolio/go-pricefeed/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.PostgresSchema,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
