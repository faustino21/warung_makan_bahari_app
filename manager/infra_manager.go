package manager

import (
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	sqlDb() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i *infra) sqlDb() *sqlx.DB {
	return i.db
}

func NewInfra(dataSourceName string) Infra {
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
