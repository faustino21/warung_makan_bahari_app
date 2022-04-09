package repository

import "Revisi_WBH/model"

type TableRepo interface {
	GetAllTable() []model.Table
	GetTable(id int) model.Table
}
