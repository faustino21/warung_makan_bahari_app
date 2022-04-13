package usecase

import (
	"Revisi_WBH/repository"
	"Revisi_WBH/util/logger"
	"fmt"
)

type CustomerUseCase interface {
	NewCustomer(name string) error
}

type customerUseCase struct {
	repo repository.CustomerRepo
}

func (c *customerUseCase) NewCustomer(name string) error {
	funcName := "CustomerUseCase.NewCustomer"

	err := c.repo.InsertCustomer(name)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return fmt.Errorf(err.Error())
	}
	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepo) CustomerUseCase {
	return &customerUseCase{
		repo: repo,
	}
}
