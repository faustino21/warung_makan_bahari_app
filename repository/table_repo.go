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
	GetStatusTable(NoTable int) (string, error)
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

func (t *tableRepo) GetStatusTable(NoTable int) (string, error) {
	var status string
	funcName := "TableRepo.GetStatusTable"

	err := t.db.Get(&status, "SELECT table_status FROM meja_makan WHERE id_table = $1", NoTable)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" :%w", err)
		return "", fmt.Errorf(funcName+" :%w", err)
	}
	return status, nil
}

func (t *tableRepo) UpdateTable(NoTable int) error {
	funcName := "TableRepo.UpdateTable"
	var status string

	err := t.db.Get(&status, "SELECT table_status FROM meja_makan WHERE id_table = $1", NoTable)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" -> Select"+" : %w", err)
		return fmt.Errorf(err.Error())
	}

	if status == "Available" {
		status = "Unavailable"
	} else {
		status = "Available"
	}

	tx := t.db.MustBegin()
	tx.MustExec("UPDATE meja_makan SET table_status = $1 WHERE id_table = $2", status, NoTable)
	err = tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" -> Update"+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewTableRepo(db *sqlx.DB) TableRepo {
	return &tableRepo{
		db: db,
	}
}
