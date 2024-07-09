// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description

package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserService,
)
