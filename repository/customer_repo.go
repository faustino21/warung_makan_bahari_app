package repository

import "Revisi_WBH/model"

type CustomerRepo interface {
	InsertCustomer(customer model.Customer)
	SelectCustomer(id int) model.Customer
}
