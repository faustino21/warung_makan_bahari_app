package repository

import (
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderDetailRepo interface {
	InsertDetail(idMenu, quantity int, IdOrder string) error
	GetTotal(idOrder string) (int, error)
}

type orderDetailRepo struct {
	db *sqlx.DB
}

func (o *orderDetailRepo) GetTotal(idOrder string) (int, error) {
	funcName := "OrderDetailRepo.GetTotal"
	var total int

	err := o.db.Get(&total,
		"SELECT SUM(quantity*mm.product_price) "+
			"FROM detail_order do2 "+
			"join menu_makanan mm on "+
			"do2.id_menu = mm.id_menu "+
			"where do2.id_order = $1",
		idOrder)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return 0, fmt.Errorf(funcName+" : %w", err)
	}
	return total, nil
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
