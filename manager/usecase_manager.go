package manager

import "Revisi_WBH/usecase"

type UseCaseManager interface {
	TableUseCase() usecase.TableUseCase
	MenuUseCase() usecase.MenuUseCase
	CustomerUseCase() usecase.CustomerUseCase
	OrderUseCase() usecase.OrderUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TableUseCase() usecase.TableUseCase {
	return usecase.NewTableUseCase(u.repo.TableRepo())
}

func (u *useCaseManager) MenuUseCase() usecase.MenuUseCase {
	return usecase.NewMenuUseCase(u.repo.MenuRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repo.CustomerRepo())
}

func (u *useCaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderRequest(u.repo.OrderRepo(), u.repo.CustomerRepo(),
		u.repo.TableRepo(), u.repo.OrderDetailRepo(), u.repo.MenuRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
