package handler

import (
	"mypos-backend/internal/model"
	"mypos-backend/internal/service"
	"mypos-backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	statusStok := c.Query("status_stok")
	if statusStok == "" {
		statusStok = "all"
	}

	categoryIdStr := c.Query("category_id")
	var categoryId int
	if categoryIdStr != "" {
		id, err := strconv.Atoi(categoryIdStr)
		if err != nil {
			response.BadRequest(c, "Format Category ID tidak valid")
			return
		}
		categoryId = id
	}

	keyword := c.Query("keyword")
	if keyword != "" {
		if len(keyword) < 3 && len(keyword) > 0 {
			response.BadRequest(c, "Keyword minimal 3 karakter")
			return
		}
	}

	if statusStok != "all" && statusStok != "ready" && statusStok != "low" {
		response.BadRequest(c, "Status stok tidak valid")
		return
	}

	products, err := h.productService.GetProduct(statusStok, categoryId, keyword)
	if err != nil {
		response.InternalServerError(c, "Gagal mendapatkan data produk:"+err.Error())
		return
	}

	if len(products) == 0 {
		response.Success(c, "Tidak ada data produk", nil)
		return
	}

	response.Success(c, "Data produk berhasil diambil", products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		response.InternalServerError(c, "Gagal mendapatkan produk:"+err.Error())
		return
	}

	if product == nil {
		response.NotFound(c, "Data produk tidak ditemukan")
		return
	}

	response.Success(c, "Data produk berhasil diambil", product)
}

func (h *ProductHandler) InsertProduct(c *gin.Context) {
	var req model.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.productService.InsertProduct(&req); err != nil {
		response.InternalServerError(c, "Gagal menambahkan data produk:"+err.Error())
		return
	}

	response.Success(c, "Berhasil menambahkan data produk", nil)
}
