// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 在这里像外部提供wire工具使用的ProviderSet

package model

import (
	"github.com/google/wire"

	"github.com/alexzhaozzzz/gin_wire_layout/internal/model/mysql"
)

var ProviderSet = wire.NewSet(
	//mysql.NewTransaction,
	//wire.Bind(new(mysqlx.Transaction), new(*mysql.Transaction)),
	mysql.NewDefaultMysql,
	//wire.Bind(new(mysqlx.IDataSource), new(*mysql.DefaultMysqlDataSource)),
	NewUserImp,
	//wire.Bind(new(repo.IUserRepo), new(*UserImp)),
)
