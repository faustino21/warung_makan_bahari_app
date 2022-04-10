package model

type Table struct {
	IdTable     int    `db:"id_table"`
	TableStatus string `db:"table_status"`
}

func NewTable(id int, status string) Table {
	return Table{
		IdTable:     id,
		TableStatus: status,
	}
}
