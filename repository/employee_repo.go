package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepo interface {
	InsertEmployee(name string) error
	GetEmployee(idEmployee int) (*model.Employee, error)
}

type employeeRepo struct {
	db *sqlx.DB
}

func (e *employeeRepo) InsertEmployee(name string) error {
	funcName := "EmployeeRepo.InsertEmployee"

	tx := e.db.MustBegin()
	tx.MustExec("INSERT INTO employee (employee_name) VALUES ($1)", name)
	err := tx.Commit()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (e *employeeRepo) GetEmployee(idEmployee int) (*model.Employee, error) {
	funcName := "CustomerRepo.GetEmployee"

	var employee model.Employee
	err := e.db.Get(&employee, "SELECT * FROM employee WHERE id_employee = $1", idEmployee)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return nil, fmt.Errorf(err.Error())
	}
	return &employee, nil
}

func NewEmployeeRepo(db *sqlx.DB) EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}
