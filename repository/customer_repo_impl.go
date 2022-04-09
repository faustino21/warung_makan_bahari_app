package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util"
	"github.com/jmoiron/sqlx"
)

type CustomerRepoImpl struct {
	customeDb *sqlx.DB
}

func (c *CustomerRepoImpl) InsertCustomer(customer model.Customer) {
	cd := c.customeDb.MustBegin()
	cd.MustExec(util.InsertCustomerQuery, customer.CustomerId, customer.CustomerName)
	cd.Commit()
}

func (c CustomerRepoImpl) SelectCustomer(id int) model.Customer {
	var customer model.Customer
	err := c.customeDb.Get(&customer, util.SelectCustomerQuery, id)
	if err != nil {
		panic(err)
	}
	return customer
}

func NewCustomerRepo(customerDb *sqlx.DB) CustomerRepo {
	return &CustomerRepoImpl{
		customeDb: customerDb,
	}
}
