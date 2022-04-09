package model

type Order struct {
	OrderId      int    `db:"order_id"`
	CustomerName string `db:"customer_name"`
	Price        int    `db:"price"`
	TableNumber  int    `db:"table_number"`
	Payment      string `db:"payment"`
}

func NewOrder(id int, customer string, price int, table int, payment string) Order {
	return Order{
		OrderId:      id,
		CustomerName: customer,
		Price:        price,
		TableNumber:  table,
		Payment:      payment,
	}
}
