package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TableRepo interface {
	GetAllTable(page int) (*[]model.Table, error)
	UpdateTable(NoTable int) error
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

func (t *tableRepo) UpdateTable(NoTable int) error {
	funcName := "TableRepo.UpdateTable"

	tx := t.db.MustBegin()
	tx.MustExec("UPDATE meja_makan SET table_status = $1 WHERE id_table = $2", "Unavailable", NoTable)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewTableRepo(db *sqlx.DB) TableRepo {
	return &tableRepo{
		db: db,
	}
}
