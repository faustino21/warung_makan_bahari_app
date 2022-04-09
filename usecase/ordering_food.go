package usecase

import (
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
)

type OrderFood interface {
	OrderingFood(orderId int, customer string, price int, table int, payment string) model.Order
	CountingOrder() int
	TotalPrice() int
}

type orderFood struct {
	repo repository.OrderRepo
}

func (o *orderFood) OrderingFood(orderId int, customer string, price int, table int, payment string) model.Order {
	newOrder := model.NewOrder(orderId, customer, price, table, payment)
	o.repo.InsertOrder(newOrder)
	return newOrder
}

func (o *orderFood) CountingOrder() int {
	return o.repo.CountOrder()
}

func (o *orderFood) TotalPrice() int {
	listOrder := o.repo.GetOrder()
	var total int
	for _, v := range listOrder {
		total += v.Price
	}
	return total
}

func NewOrderFood(repo repository.OrderRepo) OrderFood {
	return &orderFood{
		repo: repo,
	}
}
