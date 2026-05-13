package persistency

import (
	"OrderApp/common/postgresql/table"

	"gorm.io/gorm"
)

type ProductPersistency interface {
	SaveProduct(product table.Product) error
	GetProduct(id string) (*table.Product, error)
	GetPaginatedProducts(page, size int) ([]*table.Product, error)
	GetProductsByIDs(ids []string) ([]*table.Product, error)
}
type ProductPersistencyImpl struct {
	db *gorm.DB
}

func ProductAdapterImpl(db *gorm.DB) ProductPersistency {
	return &ProductPersistencyImpl{db: db}
}

func (p *ProductPersistencyImpl) SaveProduct(product table.Product) error {
	e := p.db.Create(&product).Error
	return e
}

func (p *ProductPersistencyImpl) GetProduct(id string) (*table.Product, error) {
	product := table.Product{}
	e := p.db.Where("id = ?", id).First(&product).Error
	if e != nil {
		return nil, e
	}
	return &product, nil
}

func (p *ProductPersistencyImpl) GetPaginatedProducts(page, size int) ([]*table.Product, error) {
	var products []*table.Product
	e := p.db.Limit(size).Offset((page - 1) * size).Find(&products).Error
	if e != nil {
		return nil, e
	}
	return products, nil
}

func (p *ProductPersistencyImpl) GetProductsByIDs(ids []string) ([]*table.Product, error) {
	var products []*table.Product
	e := p.db.Where("id IN (?)", ids).Find(&products).Error
	if e != nil {
		return nil, e
	}
	return products, nil
}
