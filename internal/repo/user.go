package repo

import (
	"context"
	"github.com/alexzhaozzzz/gin_wire_layout/internal/model/po"
)

// IUserRepo 用户repo接口
type IUserRepo interface {
	GetUserById(ctx context.Context, uid int64) (*po.User, error)
}

// NewUserCase ...
func NewUserCase(repo IUserRepo) *UserCase {
	return &UserCase{repo: repo}
}

// UserCase ...
type UserCase struct {
	repo IUserRepo
}

// GetPlayerById ...
func (s *UserCase) GetPlayerById(ctx context.Context) (*po.User, error) {

	return nil, nil
}
