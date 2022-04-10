package manager

import "Revisi_WBH/usecase"

type UseCaseManager interface {
	TableUseCase() usecase.TableUseCase
	MenuUseCase() usecase.MenuUseCase
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

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
