package manager

import "Revisi_WBH/repository"

type RepoManager interface {
	TableRepo() repository.TableRepo
	MenuRepo() repository.MenuRepo
	OrderDetailRepo() repository.OrderDetailRepo
	CustomerRepo() repository.CustomerRepo
	EmployeeRepo() repository.EmployeeRepo
	OrderRepo() repository.OrderRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) TableRepo() repository.TableRepo {
	return repository.NewTableRepo(r.infra.sqlDb())
}

func (r *repoManager) MenuRepo() repository.MenuRepo {
	return repository.NewMenuRepo(r.infra.sqlDb())
}

func (r *repoManager) OrderDetailRepo() repository.OrderDetailRepo {
	return repository.NewOrderDetail(r.infra.sqlDb())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.sqlDb())
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepo {
	return repository.NewEmployeeRepo(r.infra.sqlDb())
}

func (r *repoManager) OrderRepo() repository.OrderRepo {
	return repository.NewOrder(r.infra.sqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
