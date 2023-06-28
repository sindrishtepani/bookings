package dbrepo

import (
	"database/sql"

	"github.com/sindrishtepani/bookings/internal/config"
	"github.com/sindrishtepani/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DataseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewMySQLRepo(conn *sql.DB, a *config.AppConfig) {
	//return &
}
