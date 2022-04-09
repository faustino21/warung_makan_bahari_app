package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TableRepoImpl struct {
	tableDb *sqlx.DB
}

func (t *TableRepoImpl) GetAllTable() []model.Table {
	var tableList []model.Table
	err := t.tableDb.Select(&tableList, util.GetAllTableQuery)
	if err != nil {
		fmt.Println(err)
	}
	return tableList
}

func (t *TableRepoImpl) GetTable(id int) model.Table {
	var table model.Table
	err := t.tableDb.Get(&table, util.GetTableQuery, id)
	if err != nil {
		panic(err)
	}
	return table
}

func NewTableRepo(tableDb *sqlx.DB) TableRepo {
	return &TableRepoImpl{
		tableDb: tableDb,
	}
}
