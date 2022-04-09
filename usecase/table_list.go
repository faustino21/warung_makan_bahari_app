package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
)

type TableList interface {
	SearchTable() []model.Table
}

type tableList struct {
	repo repository.TableRepo
}

func (t *tableList) SearchTable() []model.Table {
	return t.repo.GetAllTable()
}

func NewTableList(repo repository.TableRepo) TableList {
	return &tableList{
		repo: repo,
	}
}
