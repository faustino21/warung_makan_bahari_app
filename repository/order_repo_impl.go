package repository

import (
	"Revisi_WBH/model"
	"Revisi_WBH/util"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderRepoImpl struct {
	orderDb *sqlx.DB
}

func (o *OrderRepoImpl) InsertOrder(newOrder model.Order) {
	od := o.orderDb.MustBegin()
	od.MustExec(util.InsertOrderQuery, newOrder.OrderId, newOrder.CustomerName, newOrder.Price, newOrder.TableNumber, newOrder.Payment)
	od.Commit()
}

func (o *OrderRepoImpl) GetOrder() []model.Order {
	var orderList []model.Order
	err := o.orderDb.Select(&orderList, util.GetAllOrderQuery)
	if err != nil {
		fmt.Println(err)
	}
	return orderList
}

func (o *OrderRepoImpl) CountOrder() int {
	var count int
	err := o.orderDb.Get(&count, util.CountOrder)
	util.IfError(err)
	return count
}

func NewOrderRepo(orderDb *sqlx.DB) OrderRepo {
	return &OrderRepoImpl{
		orderDb: orderDb,
	}
}
