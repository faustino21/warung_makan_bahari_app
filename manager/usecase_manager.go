package manager

import (
	"Revisi_WBH/usecase"
)

type UseCaseManager interface {
	CallTableList() usecase.TableList
	CallMenuList() usecase.ListMenu
	CallOrderFood() usecase.OrderFood
	CallRegisterCustomer() usecase.RegisterCustomerUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) CallOrderFood() usecase.OrderFood {
	return usecase.NewOrderFood(u.repo.OrderRepository())
}

func (u *useCaseManager) CallTableList() usecase.TableList {
	return usecase.NewTableList(u.repo.TableRepository())
}

func (u *useCaseManager) CallMenuList() usecase.ListMenu {
	return usecase.NewListMenu(u.repo.MenuRepositroy())
}

func (u *useCaseManager) CallRegisterCustomer() usecase.RegisterCustomerUseCase {
	return usecase.NewRegisterCustomer(u.repo.CustomerRepository())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
