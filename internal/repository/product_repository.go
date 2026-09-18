package repository

import (
	"database/sql"
	"mypos-backend/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProduct(statusStok string, categoryId int, keyword string) ([]model.Product, error) {
	var product []model.Product
	var args []interface{}

	query := `SELECT p.id, p.category_id, p.unit_id, p.sku, p.name, p.cost_price, p.sell_price, p.stock, p.min_stock_alert, p.is_active, COALESCE(c.name, '') as category_name, COALESCE(u.name, '') as unit_name FROM products p LEFT JOIN categories c ON p.category_id = c.id LEFT JOIN units u ON p.unit_id = u.id WHERE 1=1`

	if categoryId > 0 {
		query += ` AND p.category_id = ?`
		args = append(args, categoryId)
	}

	if keyword != "" {
		query += ` AND (p.name LIKE ? OR p.sku LIKE ?)`
		search := "%" + keyword + "%"
		args = append(args, search, search)
	}

	switch statusStok {
	case "ready":
		query += ` AND p.stock > p.min_stock_alert`
	case "low":
		query += ` AND p.stock <= p.min_stock_alert AND p.stock > 0`
	case "out":
		query += ` AND p.stock <= 0`
	}

	query += ` ORDER BY p.name ASC`

	rows, err := r.db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		err := rows.Scan(&p.ID, &p.CategoryID, &p.UnitID, &p.SKU, &p.Name, &p.CostPrice, &p.SellPrice, &p.Stock, &p.MinStockAlert, &p.IsActive, &p.CategoryName, &p.UnitName)
		if err != nil {
			return nil, err
		}

		product = append(product, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return product, err
}

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	var product model.Product

	query := `SELECT p.id, p.category_id, p.unit_id, p.sku, p.name, p.cost_price, p.sell_price, p.stock, p.min_stock_alert, p.is_active, COALESCE(c.name, '') as category_name, COALESCE(u.name, '') as unit_name FROM products p LEFT JOIN categories c ON p.category_id = c.id LEFT JOIN units u ON p.unit_id = u.id WHERE p.id = ?`
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.CategoryID, &product.UnitID, &product.SKU, &product.Name, &product.CostPrice, &product.SellPrice, &product.Stock, &product.MinStockAlert, &product.IsActive, &product.CategoryName, &product.UnitName)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &product, err
}

func (r *ProductRepository) InsertProduct(product *model.CreateProductRequest) error {
	query := `INSERT INTO products (category_id, unit_id, sku, name, cost_price, sell_price, stock, min_stock_alert) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, product.CategoryID, product.UnitID, product.SKU, product.Name, product.CostPrice, product.SellPrice, product.Stock, product.MinStockAlert)
	return err
}
