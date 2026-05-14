package persistency

import (
	"OrderApp/persistency/table"

	"gorm.io/gorm"

	"OrderApp/common/obj"
)

type ProductPersistency interface {
	SaveProduct(product table.Product) (string, error)
	GetProduct(id string) (*table.Product, error)
	GetPaginatedProducts(page, size int) ([]*table.Product, *obj.Pagination, error)
	GetProductsByIDs(ids []string) ([]*table.Product, error)
}
type ProductPersistencyImpl struct {
	db *gorm.DB
}

func ProductAdapterImpl(db *gorm.DB) ProductPersistency {
	return &ProductPersistencyImpl{db: db}
}

func (p *ProductPersistencyImpl) SaveProduct(product table.Product) (string, error) {
	e := p.db.Create(&product).Error
	if e != nil {
		return "", e
	}
	return product.ID, nil
}

func (p *ProductPersistencyImpl) GetProduct(id string) (*table.Product, error) {
	product := table.Product{}
	e := p.db.Where("id = ?", id).First(&product).Error
	if e != nil {
		return nil, e
	}
	return &product, nil
}

func (p *ProductPersistencyImpl) GetPaginatedProducts(page, size int) ([]*table.Product, *obj.Pagination, error) {
	var products []*table.Product
	e := p.db.Limit(size).Offset((page - 1) * size).Find(&products).Error
	if e != nil {
		return nil, nil, e
	}

	var count int64
	e = p.db.Model(&table.Product{}).Count(&count).Error
	if e != nil {
		return nil, nil, e
	}

	pageNums := int(count / int64(size))
	return products, &obj.Pagination{
		Page:     page,
		PageSize: size,
		PageNums: pageNums,
	}, nil
}

func (p *ProductPersistencyImpl) GetProductsByIDs(ids []string) ([]*table.Product, error) {
	var products []*table.Product
	e := p.db.Where("id IN (?)", ids).Find(&products).Error
	if e != nil {
		return nil, e
	}
	return products, nil
}
