// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 使用Google依赖注入工具wire

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/serverx"
	"github.com/google/wire"

	"github.com/alexzhaozzzz/gin_wire_layout/internal/model"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/repo"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/router"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/service"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/connect/mysqlx"
)

// initRouter 初始化router
func initRouter(ds mysqlx.IDataSource) serverx.IRouter {
	wire.Build(
		providerSet,
		router.ProviderSet,
	)
	return nil
}

var providerSet = wire.NewSet(model.ProviderSet, repo.ProviderSet, service.ProviderSet)
