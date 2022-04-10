package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
	"fmt"
	"github.com/rs/zerolog/log"
)

type MenuUseCase interface {
	GetAllMenu(page int) (*[]model.MenuMakanan, error)
}

type menuUseCase struct {
	repo repository.MenuRepo
}

func (m *menuUseCase) GetAllMenu(page int) (*[]model.MenuMakanan, error) {
	funcName := "MenuUseCase.GetAllMenu"

	menu, err := m.repo.GetAllMenu(page)
	if err != nil {
		log.Error().Msgf(funcName+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}
	return menu, nil
}

func NewMenuUseCase(repo repository.MenuRepo) MenuUseCase {
	return &menuUseCase{
		repo: repo,
	}
}
