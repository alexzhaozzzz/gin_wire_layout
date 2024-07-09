// Created on 2024-7-8 18:10:25
// @author alex
// email 1368751885@qq.com
// description

package router

import (
	"github.com/alexzhaozzzz/gin_wire_layout/internal/service"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	userSer *service.UserService
}

func NewApiRouter(userSer *service.UserService) *ApiRouter {
	return &ApiRouter{
		userSer: userSer,
	}
}

// Load 实现了server/http.go:40
func (s *ApiRouter) Load(g *gin.Engine) {
	ug := g.Group("api/v1")
	{
		ug.GET("/user", s.userSer.GetById)
	}
}
