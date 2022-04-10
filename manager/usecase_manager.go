package manager

import "Revisi_WBH/usecase"

type UseCaseManager interface {
	TableUseCase() usecase.TableUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TableUseCase() usecase.TableUseCase {
	return usecase.NewTableUseCase(u.repo.TableRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repo,
	}
}
