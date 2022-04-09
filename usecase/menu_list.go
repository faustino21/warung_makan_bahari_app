package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
)

type ListMenu interface {
	ShowMenuList() []model.MenuMakanan
}

type listMenu struct {
	repo repository.MenuMakananRepo
}

func (l *listMenu) ShowMenuList() []model.MenuMakanan {
	return l.repo.GetAllMenu()
}

func NewListMenu(repo repository.MenuMakananRepo) ListMenu {
	return &listMenu{
		repo: repo,
	}
}
