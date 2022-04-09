package manager

import "Revisi_WBH/repository"

type RepoManager interface {
	TableRepository() repository.TableRepo
	MenuRepositroy() repository.MenuMakananRepo
	OrderRepository() repository.OrderRepo
	CustomerRepository() repository.CustomerRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) TableRepository() repository.TableRepo {
	return repository.NewTableRepo(r.infra.sqlDb())
}

func (r *repoManager) MenuRepositroy() repository.MenuMakananRepo {
	return repository.NewMenuMakananRepo(r.infra.sqlDb())
}

func (r *repoManager) OrderRepository() repository.OrderRepo {
	return repository.NewOrderRepo(r.infra.sqlDb())
}

func (r *repoManager) CustomerRepository() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.sqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
