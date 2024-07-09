// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description 用户服务层

package service

import (
	"github.com/gin-gonic/gin"

	"github.com/alexzhaozzzz/gin_wire_layout/internal/repo"
)

func NewUserService(_ur repo.IUserRepo) *UserService {
	return &UserService{
		ur: _ur,
	}
}

type UserService struct {
	ur repo.IUserRepo
}

// GetById 根据用户ID查找用户
func (s *UserService) GetById(c *gin.Context) {
	info, err := s.ur.GetUserById(c, 1)
	if err != nil {
		return
	}

	c.JSON(200, info)
	return
}
