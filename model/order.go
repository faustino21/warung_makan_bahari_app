package model

type Order struct {
	OrderId     int    `db:"order_id"`
	CustomerId  int    `db:"id_customer"`
	Total       int    `db:"total"`
	TableNumber int    `db:"id_meja"`
	Payment     string `db:"payment"`
	Employee    string `db:"id_karyawan"`
}

func NewOrder(orderId, customerId, total, tableNumber int, payment, employee string) Order {
	return Order{
		OrderId:     orderId,
		CustomerId:  customerId,
		Total:       total,
		TableNumber: tableNumber,
		Payment:     payment,
		Employee:    employee,
	}
}
