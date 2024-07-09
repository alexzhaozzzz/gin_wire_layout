package model

import (
	"context"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/model/mysql"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/model/po"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/repo"
)

// NewUserImp .
func NewUserImp(db *mysql.DefaultMysqlDataSource) repo.IUserRepo {
	return &UserImp{
		db: db,
	}
}

type UserImp struct {
	db *mysql.DefaultMysqlDataSource
}

func (s *UserImp) GetUserById(ctx context.Context, uid int64) (*po.User, error) {
	info := &po.User{}
	err := s.db.Slave(ctx).Where("id", uid).First(&info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}
