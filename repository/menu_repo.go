package repository

import (
	"Revisi_WBH/model"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type MenuRepo interface {
	GetAllMenu(page int) (*[]model.MenuMakanan, error)
}

type menuRepo struct {
	db *sqlx.DB
}

func (m *menuRepo) GetAllMenu(page int) (*[]model.MenuMakanan, error) {
	funcName := "TableRepo.GetAllMenu"

	var listMenu []model.MenuMakanan
	err := m.db.Select(&listMenu, "SELECT * FROM menu_makanan ORDER BY id_menu ASC LIMIT 10 OFFSET $1 ", ((page - 1) * 10))
	if err != nil {
		log.Error().Msgf(funcName+" : %w", err)
		return nil, errors.New("Menu Not Found")
	}
	return &listMenu, nil
}

func NewMenuRepo(db *sqlx.DB) MenuRepo {
	return &menuRepo{
		db: db,
	}
}
