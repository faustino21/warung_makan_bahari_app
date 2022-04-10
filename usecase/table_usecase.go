package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
	"Revisi_WBH/util/logger"
	"fmt"
)

type TableUseCase interface {
	GetAllTable(page int) (*[]model.Table, error)
}

type tableUseCase struct {
	repo repository.TableRepo
}

func (t *tableUseCase) GetAllTable(page int) (*[]model.Table, error) {
	funcName := "TableUseCase.GetAllTable"

	table, err := t.repo.GetAllTable(page)
	if err != nil {
		logger.Log.Error().Msg("Table Not Found")
		return nil, fmt.Errorf(funcName+" : %w", err)
	}
	return table, nil
}

func NewTableUseCase(repo repository.TableRepo) TableUseCase {
	return &tableUseCase{
		repo: repo,
	}
}
