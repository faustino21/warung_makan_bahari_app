package repository

import (
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderDetailRepo interface {
	InsertDetail(idMenu, quantity int, IdOrder string) error
}

type orderDetailRepo struct {
	db *sqlx.DB
}

func (o *orderDetailRepo) InsertDetail(idMenu, quantity int, IdOrder string) error {
	funcName := "OrderDetailRepo.InsertDetail"

	tx := o.db.MustBegin()
	tx.MustExec("INSERT INTO detail_order (id_menu, quantity, id_order) VALUES ($1, $2, $3)", idMenu, quantity, IdOrder)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewOrderDetail(db *sqlx.DB) OrderDetailRepo {
	return &orderDetailRepo{
		db: db,
	}
}
