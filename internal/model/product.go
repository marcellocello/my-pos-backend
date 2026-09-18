package model

import "time"

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID            int       `json:"id"`
	CategoryID    *int      `json:"category_id"`
	CategoryName  *string   `json:"category_name,omitempty"`
	UnitID        *int      `json:"unit_id"`
	SKU           *string   `json:"sku"`
	Name          string    `json:"name"`
	CostPrice     float64   `json:"cost_price"`
	SellPrice     float64   `json:"sell_price"`
	Stock         int       `json:"stock"`
	MinStockAlert int       `json:"min_stock_alert"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	CategoryID    *int    `json:"category_id"`
	SKU           *string `json:"sku"`
	Name          string  `json:"name"`
	CostPrice     float64 `json:"cost_price"`
	SellPrice     float64 `json:"sell_price"`
	Stock         int     `json:"stock"`
	MinStockAlert int     `json:"min_stock_alert"`
}

type UpdateProductRequest struct {
	CategoryID    *int    `json:"category_id"`
	SKU           *string `json:"sku"`
	Name          string  `json:"name"`
	CostPrice     float64 `json:"cost_price"`
	SellPrice     float64 `json:"sell_price"`
	MinStockAlert int     `json:"min_stock_alert"`
	IsActive      bool    `json:"is_active"`
}

type RestockRequest struct {
	QuantityAdded int      `json:"quantity_added"`
	NewCostPrice  *float64 `json:"new_cost_price,omitempty"`
	Notes         string   `json:"notes"`
}
