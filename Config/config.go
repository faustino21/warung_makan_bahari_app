package Config

import (
	"Revisi_WBH/manager"
	"Revisi_WBH/util/logger"
	"fmt"
	"github.com/spf13/viper"
)

type ApiConfig struct {
	Url string
}

type Manager struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type DbConfig struct {
	Host     string
	User     string
	Port     string
	Password string
	Name     string
}

type Config struct {
	Manager
	ApiConfig
	DbConfig
	LogLevel string
}

func (c Config) NewConfig(path, fileName string) Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c.DbConfig = DbConfig{
		Host:     v.GetString("warungApp.db.host"),
		Port:     v.GetString("warungApp.db.port"),
		User:     v.GetString("warungApp.db.user"),
		Password: v.GetString("warungApp.db.password"),
		Name:     v.GetString("warungApp.db.name"),
	}
	c.ApiConfig = ApiConfig{Url: v.GetString("warungApp.api.url")}

	c.LogLevel = v.GetString("warungApp.api.log_level")

	return c
}

func NewConfig(path, name string) Config {
	cfg := Config{}
	cfg = cfg.NewConfig(path, name)
	logger.NewLog(cfg.LogLevel)

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	cfg.InfraManager = manager.NewInfra(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
