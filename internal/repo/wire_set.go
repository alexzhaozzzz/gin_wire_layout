// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 在这里像外部提供wire工具使用的ProviderSet

package repo

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserCase,
)
