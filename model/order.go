package model

import "time"

type Order struct {
	OrderId     string     `db:"order_id"`
	CustomerId  int        `db:"id_customer"`
	Total       int        `db:"total"`
	TableNumber int        `db:"id_meja"`
	Payment     string     `db:"payment"`
	Employee    string     `db:"id_karyawan"`
	OrderTime   *time.Time `db:"order_time"`
	PaymentTime *time.Time `db:"payment_times"`
}

func NewOrder(orderId string, customerId, total, tableNumber int, payment, employee string, timeOrder, paymentTime *time.Time) Order {
	return Order{
		OrderId:     orderId,
		CustomerId:  customerId,
		Total:       total,
		TableNumber: tableNumber,
		Payment:     payment,
		Employee:    employee,
		OrderTime:   timeOrder,
		PaymentTime: paymentTime,
	}
}
