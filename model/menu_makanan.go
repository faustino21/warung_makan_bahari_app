package model

type MenuMakanan struct {
	IdMenu       int    `db:"id_menu"`
	ProductName  string `db:"product_name"`
	ProductPrice int    `db:"product_price"`
}
