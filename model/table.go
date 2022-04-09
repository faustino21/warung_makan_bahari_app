package model

type Table struct {
	IdTable     int    `db:"id_table"`
	TableStatus string `db:"table_status"`
}

func (t *Table) GetIdTable() int {
	return t.IdTable
}

func (t *Table) GetTableStatus() string {
	return t.TableStatus
}

func NewTable(id int, status string) Table {
	return Table{
		IdTable:     id,
		TableStatus: status,
	}
}
