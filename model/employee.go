package model

type Employee struct {
	IdEmployee   int    `db:"id_employee"`
	NameEmployee string `db:"employeeName"`
}
