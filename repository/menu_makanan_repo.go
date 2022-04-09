package repository

import "Revisi_WBH/model"

type MenuMakananRepo interface {
	GetAllMenu() []model.MenuMakanan
}
