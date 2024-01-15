package repository

import (
	"database/sql"

	"github.com/GuilhermeReisOlinto/messageria_go/internal/entity"
)

type ProductRepositoryMsql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) *ProductRepositoryMsql {
	return &ProductRepositoryMsql{DB: db}
}

func (r *ProductRepositoryMsql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) VALUES(?, ?, ?)",
		product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}
