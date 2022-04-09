package Config

import (
	"Revisi_WBH/manager"
	"Revisi_WBH/util"
	"fmt"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func NewConfig() *Config {
	dataSourceName := fmt.Sprintf(util.DataSourceName, util.DbUser, util.DbPassword,
		util.DbHost, util.DbPort, util.DbName)

	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	return config
}
