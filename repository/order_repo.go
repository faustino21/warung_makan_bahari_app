package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type OrderRepo interface {
	InsertOrder(idOrder string, idCustomer, idTable, total int) error
	GetOrder(idOrder string) (*model.Order, error)
}

type orderRepo struct {
	db *sqlx.DB
}

func (o *orderRepo) InsertOrder(IdOrder string, idCustomer, idTable, total int) error {
	funcName := "OrderRepo.InsertOrder"

	timeStamp := time.Now().Local()
	tx := o.db.MustBegin()
	tx.MustExec("INSERT INTO orders (order_id, id_customer, id_meja, total, payment, order_time) VALUES ($1, $2, $3, $4, $5, $6)", IdOrder, idCustomer, idTable, total, "processing", timeStamp)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (o *orderRepo) GetOrder(idOrder string) (*model.Order, error) {
	funcName := "OrderRepo.GetOrder"

	var order model.Order
	err := o.db.Get(&order, "SELECT  * FROM orders WHERE order_id = $1", idOrder)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &order, nil
}

func NewOrder(db *sqlx.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}
