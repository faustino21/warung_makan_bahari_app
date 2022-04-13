package repository

import (
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CustomerRepo interface {
	InsertCustomer(customer string) error
	CountCustomer() (int, error)
}

type customerRepo struct {
	db *sqlx.DB
}

func (c *customerRepo) CountCustomer() (int, error) {
	funcName := "CustomerRepo.CountCustomer"
	var result int

	err := c.db.Get(&result, "SELECT COUNT(customer_id) FROM customers")
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return 0, fmt.Errorf(funcName+" :%w", err)
	}
	return result, nil
}

func (c *customerRepo) InsertCustomer(customer string) error {
	funcName := "CustomerRepo.InsertCustomer"

	tx := c.db.MustBegin()
	tx.MustExec("INSERT INTO customers (customer_name) VALUES ($1)", customer)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepo {
	return &customerRepo{
		db: db,
	}
}
