// author: xmgtony
// date: 2023-06-29 14:47
// version:

package mysql

import (
	"context"
	"github.com/alexzhaozzzz/gin_wire_layout/pkg/connect/mysqlx"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// DefaultMysqlDataSource 默认mysql数据源实现
type DefaultMysqlDataSource struct {
	master *gorm.DB // 定义私有属性，用来持有主库链接，防止每次创建，创建后直接返回该变量。
	slave  *gorm.DB // 同上，从库链接
}

func (s *DefaultMysqlDataSource) Master(ctx context.Context) *gorm.DB {
	// 事物, 根据事物的key取出tx
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if s.master == nil {
		panic("The [master] connection is nil, Please initialize it first.")
	}
	return s.master
}

func (s *DefaultMysqlDataSource) Slave(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if s.slave == nil {
		panic("The [slave] connection is nil, Please initialize it first.")
	}
	return s.slave
}

func (s *DefaultMysqlDataSource) Close() {
	// 关闭主库链接
	if s.master != nil {
		m, err := s.master.DB()
		if err != nil {
			_ = m.Close()
		}
	}
	// 关闭从库链接
	if s.slave != nil {
		_s, err := s.slave.DB()
		if err != nil {
			_ = _s.Close()
		}
	}
}

func NewDefaultMysql() *DefaultMysqlDataSource {
	master := mysqlx.NewMysqlConn(
		viper.GetString("mysql-dbs.merchant.user"),
		viper.GetString("mysql-dbs.merchant.password"),
		viper.GetString("mysql-dbs.merchant.host"),
		viper.GetString("mysql-dbs.merchant.port"),
		viper.GetString("mysql-dbs.merchant.database"),
		viper.GetInt("mysql-dbs.merchant.max_open"),
		viper.GetInt("mysql-dbs.merchant.max_idle"))

	slave := mysqlx.NewMysqlConn(
		viper.GetString("mysql-dbs.merchant.user"),
		viper.GetString("mysql-dbs.merchant.password"),
		viper.GetString("mysql-dbs.merchant.host"),
		viper.GetString("mysql-dbs.merchant.port"),
		viper.GetString("mysql-dbs.merchant.database"),
		viper.GetInt("mysql-dbs.merchant.max_open"),
		viper.GetInt("mysql-dbs.merchant.max_idle"))

	return &DefaultMysqlDataSource{
		master: master,
		slave:  slave,
	}
}
