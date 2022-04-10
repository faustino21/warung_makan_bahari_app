package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util/logger"
	uuid2 "Revisi_WBH/util/uuid"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type OrderRepo interface {
	InsertOrder(idCustomer, idTable, total int, idEmployee string) error
	GetOrder(idOrder string) error
}

type orderRepo struct {
	db *sqlx.DB
}

func (o *orderRepo) InsertOrder(idCustomer, idTable, total int, idEmployee string) error {
	funcName := "OrderRepo.InsertOrder"

	timeStamp := time.Now().Local()
	uuid := uuid2.GetUuid()

	tx := o.db.MustBegin()
	tx.MustExec("INSERT INTO orders (order_id, id_customer, id_meja, total, payment, id_karyawan, time_order) VALUES ($1, $2, $3, $4, $5, $6, $7)", uuid, idCustomer, idTable, total, "processing", idEmployee, timeStamp)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (o *orderRepo) GetOrder(idOrder string) error {
	funcName := "OrderRepo.GetOrder"

	var order model.Order
	err := o.db.Get(&order, "SELECT  * FROM orders WHERE order_id = $1", idOrder)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewOrder(db *sqlx.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}
