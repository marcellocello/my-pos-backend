package model

import "time"

type PaymentMethod string

const (
	PaymentMethodCash PaymentMethod = "cash"
	PaymentMethodQRIS PaymentMethod = "qris"
)

type Transaction struct {
	ID            int               `json:"id"`
	InvoiceNumber string            `json:"invoice_number"`
	UserID        int               `json:"user_id"`
	CashierName   string            `json:"cashier_name,omitempty"`
	TotalAmount   float64           `json:"total_amount"`
	PaidAmount    float64           `json:"paid_amount"`
	ChangeAmount  float64           `json:"change_amount"`
	PaymentMethod PaymentMethod     `json:"payment_method"`
	Items         []TransactionItem `json:"items,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
}

type TransactionItem struct {
	ID            int     `json:"id"`
	TransactionID int     `json:"transaction_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name,omitempty"`
	Quantity      int     `json:"quantity"`
	UnitCostPrice float64 `json:"unit_cost_price"`
	UnitSellPrice float64 `json:"unit_sell_price"`
	Subtotal      float64 `json:"subtotal"`
}

type CheckoutItemRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items         []CheckoutItemRequest `json:"items"`
	PaymentMethod PaymentMethod         `json:"payment_method"`
	PaidAmount    float64               `json:"paid_amount"`
}

type StockLogType string

const (
	StockLogSale       StockLogType = "sale"
	StockLogRestock    StockLogType = "restock"
	StockLogAdjustment StockLogType = "adjustment"
)

type StockLog struct {
	ID             int          `json:"id"`
	ProductID      int          `json:"product_id"`
	ProductName    string       `json:"product_name,omitempty"`
	QuantityChange int          `json:"quantity_change"`
	Type           StockLogType `json:"type"`
	ReferenceID    *string      `json:"reference_id"`
	Notes          *string      `json:"notes"`
	CreatedAt      time.Time    `json:"created_at"`
}
