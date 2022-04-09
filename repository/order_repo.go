package repository

import "Revisi_WBH/model"

type OrderRepo interface {
	GetOrder() []model.Order
	InsertOrder(newOrder model.Order)
	CountOrder() int
}
