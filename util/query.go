package util

const (
	GetAllTableQuery    = "SELECT * FROM meja_makan WHERE table_status = 'Available' ORDER BY id_table ASC"
	GetAllMenuQuery     = "SELECT * FROM menu_makanan ORDER BY id_menu"
	InsertOrderQuery    = "INSERT INTO orders (order_id, customer_name, price, table_number, payment) VALUES ($1, $2, $3, $4, $5)   "
	GetAllOrderQuery    = "SELECT * FROM orders ORDER BY order_id "
	SelectMenuQuery     = "SELECT * FROM menu_makanan WHERE id_menu = $1"
	InsertCustomerQuery = "INSERT INTO customers (customer_id, customer_name) VALUES ($1, $2)"
	SelectCustomerQuery = "SELECT * FROM customers WHERE customer_id = $1"
	GetTableQuery       = "SELECT * FROM meja_makan WHERE id_table = $1"
	CountOrder          = "SELECT COUNT(order_id) FROM orders"
)
