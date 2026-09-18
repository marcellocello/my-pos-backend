package handler

import (
	"mypos-backend/internal/model"
	"mypos-backend/internal/service"
	"mypos-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnitHandler struct {
	unitService *service.UnitService
}

func NewUnitHandler(unitService *service.UnitService) *UnitHandler {
	return &UnitHandler{unitService: unitService}
}

func (h *UnitHandler) GetAllUnits(c *gin.Context) {
	units, err := h.unitService.GetAllUnits()
	if err != nil {
		response.BadRequest(c, "Gagal mengambil data units:"+err.Error())
	}

	if len(units) == 0 {
		response.Success(c, "Tidak ada data units", nil)
	}

	response.Success(c, "Data units berhasil diambil", units)
}

func (h *UnitHandler) InsertUnit(c *gin.Context) {
	var req model.CreateUnitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.unitService.InsertUnit(&req); err != nil {
		response.InternalServerError(c, "Gagal menambahkan unit:"+err.Error())
		return
	}

	response.Created(c, "Berhasil menambahkan unit", nil)
}

func (h *UnitHandler) UpdateUnit(c *gin.Context) {
	var req model.UpdateUnitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.unitService.UpdateUnit(&req); err != nil {
		response.InternalServerError(c, "Gagal update data unit:"+err.Error())
		return
	}

	response.Success(c, "Berhasil update data unit", nil)
}

func (h *UnitHandler) DeleteUnit(c *gin.Context) {
	unitId := c.Param("id")

	if unitId == "" {
		response.BadRequest(c, "Unit ID harus diisi")
		return
	}

	unitIdConverted, err := strconv.Atoi(unitId)
	if err != nil {
		response.InternalServerError(c, "Gagal menghapus data unit:"+err.Error())
		return
	}

	err = h.unitService.DeleteUnit(unitIdConverted)
	if err != nil {
		response.InternalServerError(c, "Gagal menghapus data unit:"+err.Error())
		return
	}

	response.Success(c, "Berhasil menghapus data unit", nil)
}

func (h *UnitHandler) GetUnitByID(c *gin.Context) {
	unitId := c.Param("id")

	if unitId == "" {
		response.BadRequest(c, "Unit ID harus diisi")
		return
	}

	unitIdConverted, err := strconv.Atoi(unitId)
	if err != nil {
		response.InternalServerError(c, "Gagal mendapatkan data unit:"+err.Error())
		return
	}

	unit, err := h.unitService.GetUnitByID(unitIdConverted)
	if err != nil {
		response.InternalServerError(c, "Gagal mendapatkan data unit:"+err.Error())
		return
	}

	if unit == nil {
		response.NotFound(c, "Data unit dengan ID '"+unitId+"' tidak ditemukan")
		return
	}

	response.Success(c, "Berhasil mendapatkan data unit", unit)
}
