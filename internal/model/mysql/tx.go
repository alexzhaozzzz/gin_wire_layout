// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 事物控制接口

package mysql

import (
	"context"
	"gorm.io/gorm"

	"github.com/alexzhaozzzz/gin_wire_layout/pkg/connect/mysqlx"
)

type contextTxKey struct{}

// Transaction 事物默认实现
type Transaction struct {
	ds mysqlx.IDataSource
}

func NewTransaction(_ds mysqlx.IDataSource) *Transaction {
	return &Transaction{ds: _ds}
}

func (t *Transaction) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.ds.Master(ctx).Transaction(func(tx *gorm.DB) error {
		withValue := context.WithValue(ctx, contextTxKey{}, tx)
		return fn(withValue)
	})
}
