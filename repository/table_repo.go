package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TableRepo interface {
	GetAllTable(page int) (*[]model.Table, error)
}
type tableRepo struct {
	db *sqlx.DB
}

func (t *tableRepo) GetAllTable(page int) (*[]model.Table, error) {
	funcName := "repository.GetAllTable"

	var tableList []model.Table
	err := t.db.Select(&tableList, "SELECT * FROM meja_makan ORDER BY id_table ASC LIMIT 5 OFFSET $1", ((page - 1) * 5))
	if err != nil {
		logger.Log.Error().Msg(err.Error())
		return nil, fmt.Errorf(funcName+" : %w", err)
	}
	return &tableList, nil
}

func NewTableRepo(db *sqlx.DB) TableRepo {
	return &tableRepo{
		db: db,
	}
}
