package model

type DetailOrder struct {
	IdDetail string `db:"id_detail"`
	IdMenu   int    `db:"id_menu"`
	Quantity int    `db:"quantity"`
}
