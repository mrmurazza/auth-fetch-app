package handler

import (
	"authapp/domain/user"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	userSvc user.Service
}

func NewApiHandler(userSvc user.Service) *ApiHandler {
	return &ApiHandler{
		userSvc: userSvc,
	}
}

func (h *ApiHandler) Login(c *gin.Context) {
}

func (h *ApiHandler) CheckAuth(c *gin.Context) {
}

func (h *ApiHandler) CreateUser(c *gin.Context) {
}
