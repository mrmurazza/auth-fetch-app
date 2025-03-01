package handler

import (
	"authapp/domain/user"
	"authapp/dto/request"
	"net/http"

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
	var req request.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	u, signedToken, err := h.userSvc.Login(req.Phonenumber, req.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not registered"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   signedToken,
		"message": "success",
	})
}

func (h *ApiHandler) CheckAuth(c *gin.Context) {
	userInfo, ok := c.Get("userInfo")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userinfo not filled"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    userInfo,
		"message": "success",
	})
}

func (h *ApiHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = req.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	u, err := h.userSvc.CreateUserIfNotAny(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"password": u.Password,
		},
		"message": "success",
	})
}
