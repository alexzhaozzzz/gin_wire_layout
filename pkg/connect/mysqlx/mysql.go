// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 配置mysql链接

package mysqlx

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// IDataSource 定义数据库数据源接口，按照业务需求可以返回主库链接Master和从库链接Slave
type IDataSource interface {
	Master(ctx context.Context) *gorm.DB
	Slave(ctx context.Context) *gorm.DB
	Close()
}

// NewMysqlConn 创建Mysql链接
func NewMysqlConn(user, password, host, port, dbname string, maxPoolSize, maxIdle int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	if err != nil {
		panic(fmt.Sprintf("mysql open error: %s", err))
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("mysql db error: %s", err))
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(maxPoolSize)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Duration(30) * time.Second)

	return db
}
