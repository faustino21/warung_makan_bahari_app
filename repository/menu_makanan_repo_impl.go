package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MenuMakananRepoImpl struct {
	menuMakananDb *sqlx.DB
}

func (m *MenuMakananRepoImpl) GetAllMenu() []model.MenuMakanan {
	var menuMakanan []model.MenuMakanan
	err := m.menuMakananDb.Select(&menuMakanan, util.GetAllMenuQuery)
	if err != nil {
		panic(err)
	}
	return menuMakanan
}

func (m *MenuMakananRepoImpl) SelectMenu(foodId int) model.MenuMakanan {
	var food model.MenuMakanan
	err := m.menuMakananDb.Get(&food, util.SelectMenuQuery, foodId)
	if err != nil {
		fmt.Println(err)
	}
	return food
}

func NewMenuMakananRepo(menuDb *sqlx.DB) MenuMakananRepo {
	return &MenuMakananRepoImpl{
		menuMakananDb: menuDb,
	}
}
