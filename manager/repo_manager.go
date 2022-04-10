package manager

import "Revisi_WBH/repository"

type RepoManager interface {
	TableRepo() repository.TableRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) TableRepo() repository.TableRepo {
	return repository.NewTableRepo(r.infra.sqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
