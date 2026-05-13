package persistence

import (
	"OrderApp/common/postgresql/table"
	"OrderApp/service/inventory/application/domain/model"

	"gorm.io/gorm"
)

type ProductAdapter struct {
	db *gorm.DB
}

func ProductAdapterImpl(db *gorm.DB) *ProductAdapter {
	return &ProductAdapter{db: db}
}

func (p *ProductAdapter) SaveProduct(product *model.Product) error {
	salePrice := ""
	if product.GetSalePrice() != nil {
		salePrice = (*product.GetSalePrice()).String()
	}

	tProduct := table.Product{
		Name:      product.GetName(),
		Price:     product.GetPrice().String(),
		SalePrice: &salePrice,
	}

	e := p.db.Create(&tProduct).Error
	return e
}

func (p *ProductAdapter) GetProduct(id string) (*model.Product, error) {
	persistenceProduct := table.Product{}
	e := p.db.Where("id = ?", id).First(&persistenceProduct).Error
	if e != nil {
		return nil, e
	}

	product, e := MapProduct(persistenceProduct)
	if e != nil {
		return nil, e
	}
	return product, nil
}

func (p *ProductAdapter) GetPaginatedProducts(page, size int) ([]*model.Product, error) {
	var persistenceProducts []table.Product
	e := p.db.Limit(size).Offset((page - 1) * size).Find(&persistenceProducts).Error
	if e != nil {
		return nil, e
	}

	products := make([]*model.Product, len(persistenceProducts))
	for i, persistenceProduct := range persistenceProducts {
		products[i], e = MapProduct(persistenceProduct)
		if e != nil {
			return nil, e
		}
	}

	return products, nil
}

func (p *ProductAdapter) GetProductsByIDs(ids []string) ([]*model.Product, error) {
	var persistenceProducts []table.Product
	e := p.db.Where("id IN (?)", ids).Find(&persistenceProducts).Error
	if e != nil {
		return nil, e
	}

	products := make([]*model.Product, len(persistenceProducts))
	for i, persistenceProduct := range persistenceProducts {
		products[i], e = MapProduct(persistenceProduct)
		if e != nil {
			return nil, e
		}
	}
	return products, nil
}
