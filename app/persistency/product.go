package persistency

import (
	"OrderApp/common/msg"
	"OrderApp/persistency/table"
	"errors"

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
	e := p.db.
		Limit(size).
		Offset((page - 1) * size).
		Find(&products).
		Order("createdAt DESC").
		Error
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

	isIdNotFound := len(products) != len(ids)
	if isIdNotFound {
		return nil, errors.New(msg.ProductNotFound)
	}

	// restore product order
	findById := func(id string, products []*table.Product) (*table.Product, error) {
		for _, prod := range products {
			if prod.ID == id {
				return prod, nil
			}
		}
		return nil, errors.New(msg.ProductNotFound)
	}

	restoredProducts := make([]*table.Product, len(ids))
	for i, _ := range restoredProducts {
		prod, e := findById(ids[i], products)
		if e != nil {
			return nil, e
		}
		restoredProducts[i] = prod
	}

	return restoredProducts, nil
}
