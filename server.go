package main

import (
	"Revisi_WBH/Config"
	"Revisi_WBH/Delivery/api"
	"Revisi_WBH/manager"
	"Revisi_WBH/util/logger"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	r              *gin.Engine
	cfg            Config.Config
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func (a *appServer) initHandler() {
	a.v1()
}

func (a *appServer) v1() {
	orderRoute := a.r.Group("/api")
	api.TableApiRoute(orderRoute, a.cfg.UseCaseManager.TableUseCase())
	api.MenuApiRoute(orderRoute, a.cfg.UseCaseManager.MenuUseCase())
}

func (a appServer) Run() {
	a.initHandler()
	err := a.r.Run(a.cfg.ApiConfig.Url)
	if err != nil {
		logger.Log.Fatal().Msg("Server failed to run")
	}
}

func Server() AppServer {
	r := gin.Default()
	c := Config.NewConfig(".", "config")

	return &appServer{
		r:   r,
		cfg: c,
	}
}
