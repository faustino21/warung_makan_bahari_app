package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
)

type RegisterCustomerUseCase interface {
	RegisterCustomer(CustomerId int, CustomerName string) model.Customer
}

type registerCustomerUseCase struct {
	repo repository.CustomerRepo
}

func (r registerCustomerUseCase) RegisterCustomer(customerId int, customerName string) model.Customer {
	newCustomer := model.NewCustomer(customerId, customerName)
	r.repo.InsertCustomer(newCustomer)
	return newCustomer
}

func NewRegisterCustomer(repo repository.CustomerRepo) RegisterCustomerUseCase {
	return &registerCustomerUseCase{
		repo: repo,
	}
}
