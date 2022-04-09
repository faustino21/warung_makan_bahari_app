package model

type Customer struct {
	CustomerId   int    `db:"customer_id"`
	CustomerName string `db:"customer_name"`
}

func NewCustomer(id int, name string) Customer {
	return Customer{
		CustomerId:   id,
		CustomerName: name,
	}
}
