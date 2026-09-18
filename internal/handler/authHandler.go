package handler

import (
	"mypos-backend/internal/model"
	"mypos-backend/internal/service"
	"mypos-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, token, err := h.authService.Login(&req)
	if err != nil {
		response.Unauthorized(c, "Username atau Password salah")
		return
	}

	loginResponse := model.LoginResponse{
		Token: token,
		User:  *user,
	}

	response.Success(c, "Login berhasil", loginResponse)
}
