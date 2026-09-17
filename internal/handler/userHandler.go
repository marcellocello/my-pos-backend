package handler

import (
	"mypos-backend/internal/model"
	"mypos-backend/internal/service"
	"mypos-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req model.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.userService.Register(&req); err != nil {
		response.InternalServerError(c, "Gagal mendaftarkan user:"+err.Error())
		return
	}

	response.Created(c, "User berhasil didaftarkan", nil)
}

func (h *UserHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		response.BadRequest(c, "Username wajib diisi")
		return
	}

	user, err := h.userService.GetByUsername(username)
	if err != nil {
		response.InternalServerError(c, "Gagal mengambil data user:"+err.Error())
		return
	}

	if user == nil {
		response.NotFound(c, "User dengan username '"+username+"' tidak ditemukan")
		return
	}

	response.Success(c, "Berhasil mendapatkan data user", user)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req model.ChangePasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Validasi gagal:"+err.Error())
		return
	}

	if err := h.userService.ChangePassword(&req); err != nil {
		if err.Error() == "password lama salah" || err.Error() == "user tidak ditemukan" {
			response.BadRequest(c, err.Error())
			return
		}

		response.InternalServerError(c, "Gagal mengubah password:"+err.Error())
		return
	}

	response.Success(c, "Password berhasil diperbarui", nil)
}
