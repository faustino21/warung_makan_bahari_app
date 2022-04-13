package usecase

import (
	"Revisi_WBH/Delivery/httpReq"
	"Revisi_WBH/model"
	"Revisi_WBH/repository"
	"Revisi_WBH/util/logger"
	uuid2 "Revisi_WBH/util/uuid"
	"fmt"
)

type OrderUseCase interface {
	OrderProduct(order httpReq.OrderReq) (*model.Order, error)
}

type orderUseCase struct {
	orderRepo    repository.OrderRepo
	customerRepo repository.CustomerRepo
	tableRepo    repository.TableRepo
	detailRepo   repository.OrderDetailRepo
	menuRepo     repository.MenuRepo
}

func (o *orderUseCase) OrderProduct(order httpReq.OrderReq) (*model.Order, error) {
	funcName := "OrderUseCase.OrderProduct"

	//INSERT CUSTOMER
	err := o.customerRepo.InsertCustomer(order.Customer)
	if err != nil {
		logger.Log.Error().Msg(err.Error())
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	//CHECKING TABLE
	status, err := o.tableRepo.GetStatusTable(order.Table)
	if status != "Available" {
		logger.Log.Error().Msgf(funcName + "")
		return nil, fmt.Errorf("table not available")
	}
	if err != nil {
		logger.Log.Error().Msgf(funcName+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	//INSERT ORDER TABLE
	err = o.tableRepo.UpdateTable(order.Table)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" update error"+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	//INSERT ORDER DETAIL
	uuid := uuid2.GetUuid()
	for _, v := range order.ListMenu {
		err = o.detailRepo.InsertDetail(v.IdMenu, v.Quantity, uuid)
		if err != nil {
			logger.Log.Error().Msgf(funcName+"insert detail error"+" : %w", err)
			return nil, fmt.Errorf(funcName+" : %w", err)
		}
	}

	//INSERT ORDER
	res, err := o.customerRepo.CountCustomer()
	if err != nil {
		logger.Log.Error().Msgf(funcName+" Count Customer Error"+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	total, err := o.detailRepo.GetTotal(uuid)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" Get Total Error"+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	err = o.orderRepo.InsertOrder(uuid, res, order.Table, total)
	if err != nil {
		logger.Log.Error().Msgf(funcName+"Insert Order Error"+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	result, err := o.orderRepo.GetOrder(uuid)
	if err != nil {
		logger.Log.Error().Msgf(funcName+" Get Order Error"+" : %w", err)
		return nil, fmt.Errorf(funcName+" : %w", err)
	}

	return result, nil
}

func NewOrderRequest(order repository.OrderRepo, customer repository.CustomerRepo, table repository.TableRepo,
	detail repository.OrderDetailRepo, menu repository.MenuRepo) OrderUseCase {
	return &orderUseCase{
		orderRepo:    order,
		tableRepo:    table,
		customerRepo: customer,
		detailRepo:   detail,
		menuRepo:     menu,
	}
}
