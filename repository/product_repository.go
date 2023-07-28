package repository

import (
	"database/sql"

	"github.com/septian03yogi/enigmalaundryinc/model"
	"github.com/septian03yogi/enigmalaundryinc/model/dto"
	"github.com/septian03yogi/enigmalaundryinc/utils/common"
)

type ProductRepository interface {
	BaseRepository[model.Product]
	BaseRepositoryPaging[model.Product]
}

type productRepository struct {
	db *sql.DB
}

// Create Implements ProductRepository
func (p *productRepository) Create(payload model.Product) error {
	_, err := p.db.Exec("INSERT INTO product (id, name, price, uom_id) VALUES ($1, $2, $3, $4)", payload.Id, payload.Name, payload.Price, payload.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

// List implements productrepository
func (p *productRepository) List() ([]model.Product, error) {
	rows, err := p.db.Query("SELECT p.id, p.name, p.price, u.id, u.name FROM product p INNER JOIN uom u ON u.id = p.uom_id")
	if err != nil {
		return nil, err
	}

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Get implements productrepository
func (p *productRepository) Get(id string) (model.Product, error) {
	var product model.Product
	row := p.db.QueryRow("SELECT p.id, p.name, p.price, u.id, u.name FROM product p INNER JOIN uom u ON u.id=p.uom_id WHERE p.id = $1", id)
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// Update implementss productrepository
func (p *productRepository) Update(payload model.Product) error {
	_, err := p.db.Exec("UPDATE product SET name = $2, price = $3, uom_id=$4 WHERE id=$1", payload.Id, payload.Name, payload.Price, payload.Uom.Id)
	if err != nil {
		return err
	}
	return nil
}

// Delete Implements ProductRepository
func (p *productRepository) Delete(id string) error {
	_, err := p.db.Exec("DELETE FROM product WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

// Paging implements productrepository
func (p *productRepository) Paging(requestPaging dto.PaginationParam) ([]model.Product, dto.Paging, error) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)

	rows, err := p.db.Query("SELECT p.id, p.name, p.price, u.id, u.name FROM product p INNER JOIN uom u ON u.id=p.uom_id LIMIT $1 OFFSET$2", paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		products = append(products, product)
	}
	//count product
	var totalRows int
	row := p.db.QueryRow("SELECT COUNT * FROM product")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return products, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
