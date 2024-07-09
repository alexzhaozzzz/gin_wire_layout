// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description

package router

import (
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/serverx"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewApiRouter,
	wire.Bind(new(serverx.IRouter), new(*ApiRouter)),
)
